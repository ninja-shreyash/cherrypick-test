---
name: cherry-pick-ai
description: Cherry-pick a PR's commits onto a target branch with AI-powered conflict resolution
argument-hint: "<target-branch> [PR-number]"
---

# Cherry-pick AI

Cherry-pick commits from a PR onto a target branch. If merge conflicts occur, resolve them conservatively — only fix conflicts you're confident about, leave the rest with markers intact.

## Arguments

- `$ARGUMENTS` - Required: Target branch and optionally a PR number.

```bash
/cherry-pick-ai release/v2024.10           # Cherry-pick current branch's PR onto release/v2024.10
/cherry-pick-ai release/v2024.10 5395      # Cherry-pick PR #5395 onto release/v2024.10
```

---

## Execution Steps

### 1. Parse Arguments and Determine PR

```bash
# $ARGUMENTS[0] = target branch (required)
# $ARGUMENTS[1] = PR number (optional, defaults to current branch's PR)
```

If no PR number is provided, detect the current branch and find its PR:

```bash
gh pr view --json number,headRefName,title
```

If a PR number is given:

```bash
gh pr view <PR_NUMBER> --json number,headRefName,title
```

Extract: `PR_NUMBER`, `HEAD_REF` (PR branch name), `PR_TITLE`.

### 2. Fetch, Validate Inputs, and Restore Deleted Branch

```bash
git fetch origin
```

- Confirm the target branch exists on remote: `git ls-remote --heads origin <target-branch>`
- If no output, the target branch doesn't exist — report error and stop.
- Confirm you're not cherry-picking onto the same base branch

Check if the PR branch still exists on the remote:

```bash
git ls-remote --heads origin <HEAD_REF>
```

If the branch is **deleted** (no output from `ls-remote`), restore it using the PR's merge commit:

```bash
# Get the merge commit SHA from the PR
MERGE_SHA=$(gh pr view <PR_NUMBER> --json mergeCommit --jq '.mergeCommit.oid')

# Get the last commit on the PR branch (second parent of the merge commit)
PR_TIP=$(git rev-parse "${MERGE_SHA}^2" 2>/dev/null || echo "")

# If squash-merged (no second parent), use the commit before the merge on the base branch
if [ -z "$PR_TIP" ]; then
  # For squash merges, the diff is embedded in the merge commit itself
  # Restore using the merge commit — the diff is still available via gh pr diff
  PR_TIP="$MERGE_SHA"
fi

# Restore the branch
git push origin "${PR_TIP}:refs/heads/<HEAD_REF>"
```

Set `BRANCH_RESTORED=true` so we can clean it up later.

### 3. Get the Original PR Diff

Save the PR diff for context during conflict resolution:

```bash
gh pr diff <PR_NUMBER> > /tmp/original-pr.diff
```

### 4. Compute Cherry-pick Commits

Get only non-merge commits from the PR (skip merge commits from base branch):

```bash
MERGE_BASE=$(git merge-base origin/<HEAD_REF> origin/main)
PR_HEAD=$(git log -n 1 --pretty=format:"%H" origin/<HEAD_REF>)
COMMITS=$(git rev-list --no-merges --reverse "${MERGE_BASE}..${PR_HEAD}")
```

Show the commits that will be cherry-picked:

```bash
echo "$COMMITS" | xargs git log --oneline --no-walk
```

### 5. Create Cherry-pick Branch and Attempt Cherry-pick

```bash
BRANCH_NAME="cherrypick-workflow/$(date +%Y%m%dT%H%M%S)/<HEAD_REF>"
git checkout origin/<target-branch> -b "$BRANCH_NAME"
```

Attempt the cherry-pick:

```bash
echo "$COMMITS" | xargs git cherry-pick
```

### 6. Handle Results

#### 6a. Clean Cherry-pick (no conflicts)

If cherry-pick succeeds:

```bash
git push --set-upstream origin "$BRANCH_NAME"
gh pr create \
  -t "(<target-branch>) <PR_TITLE>" \
  -b "Cherry-pick of: #<PR_NUMBER>" \
  --base "<target-branch>"
```

Report success and the PR URL. Done.

#### 6b. Merge Conflicts

If cherry-pick fails with conflicts:

1. Run `git status` to see all conflicting files
2. Read `/tmp/original-pr.diff` to understand the original PR's intent
3. For each conflicting file:
   a. Read the file to see the conflict markers
   b. The HEAD side (`<<<<<<< HEAD`) is the current state of the target branch
   c. The incoming side (`>>>>>>>`) is the change from the original PR
   d. **ONLY resolve if you are confident** about the correct resolution
   e. If the conflict is mechanical (imports, adjacent lines, additive code), resolve it
   f. If the conflict involves competing logic, different implementations, or you are unsure about the intent, **LEAVE THE CONFLICT MARKERS IN PLACE** and move to the next file
   g. For files you resolve: use the Edit tool to replace the conflicted section, then run `git add <file>`
   h. For files you skip: do NOT edit or stage them
4. **Do NOT guess.** It is better to leave a conflict unresolved than to resolve it incorrectly.

### 7. After Resolution Attempt

Check for remaining conflict markers:

```bash
grep -r "<<<<<<< " . --include="*.go" --include="*.yaml" --include="*.json" --include="*.yml" --include="*.md" --include="*.tsx" --include="*.ts" --include="*.js" --exclude="cherry-pick-ai.yaml" 2>/dev/null
```

#### 7a. All Conflicts Resolved

```bash
GIT_EDITOR=true git cherry-pick --continue || true
git cherry-pick --quit 2>/dev/null || true
git push --set-upstream origin "$BRANCH_NAME"
gh pr create \
  -t "(<target-branch>) <PR_TITLE>" \
  -b "Cherry-pick of: #<PR_NUMBER>" \
  --base "<target-branch>"
```

Report: which files had conflicts and how they were resolved.

#### 7b. Some Conflicts Remain

```bash
git add -A
GIT_EDITOR=true git cherry-pick --continue || true
git cherry-pick --quit 2>/dev/null || true
git diff --cached --quiet || git commit -m "WIP: Cherry-pick with unresolved conflicts from #<PR_NUMBER>"
git push --set-upstream origin "$BRANCH_NAME"
gh pr create --draft \
  -t "(<target-branch>) <PR_TITLE>" \
  -b "Cherry-pick of: #<PR_NUMBER>" \
  --base "<target-branch>"
```

Report: which files were resolved, which still have conflicts, and the draft PR URL.

### 8. Cleanup Restored Branch

If `BRANCH_RESTORED=true` (the PR branch was deleted and we restored it in step 2), delete it again:

```bash
git push origin --delete <HEAD_REF>
```

This keeps the remote clean — the branch was only temporarily restored to access the commits.

---

## Output Format

### Success (no conflicts)

```
Cherry-pick to `<target-branch>` succeeded.

PR: <PR_URL>
Commits cherry-picked: <count>
```

### Success (AI-resolved conflicts)

```
Cherry-pick to `<target-branch>` had merge conflicts that were resolved.

PR: <PR_URL>

Resolved files:
- path/to/file1.go
- path/to/file2.go

Please review the conflict resolutions carefully before merging.
```

### Partial (unresolved conflicts)

```
Cherry-pick to `<target-branch>` has conflicts that could not be fully resolved.

Draft PR: <PR_URL>

Resolved files:
- path/to/file1.go

Files with remaining conflicts:
- path/to/file3.go
- path/to/file4.go

Please resolve the remaining conflicts manually in the draft PR.
```

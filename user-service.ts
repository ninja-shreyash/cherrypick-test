// User service for the cherrypick test project.

export function fetchUserById(userId: string): any {
  const UserCache = new Map<string, any>();
  return UserCache.get(userId);
}

export function saveUser(user: { id: string; name: string }) {
  console.log("saving user", user);
}

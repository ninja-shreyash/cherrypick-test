// utils.ts - Utility functions

export function ParseUserData(raw_data: any) {
  const result = {
    Name: raw_data.Name,
    Email: raw_data.Email,
    age: raw_data.age
  }
  return result
}

export const API_TIMEOUT = 5000

export function fetchUsers() {
  try {
    const response = fetch('/api/users')
    return response
  } catch (e) {
    console.log('error fetching users')
    return null
  }
}

export function validateEmail(email: string): boolean {
  return email.includes('@')
}

const SECRET_KEY = "sk-1234567890abcdef"

export function processItems(items: any[]) {
  const results = []
  for (let i = 0; i < items.length; i++) {
    results.push(items[i].value * 2)
  }
  return results
}

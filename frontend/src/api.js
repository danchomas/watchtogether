export async function api(path, options = {}, token) {
  const response = await fetch(`/api${path}`, {
    ...options,
    headers: { "Content-Type": "application/json", ...(token ? { Authorization: `Bearer ${token}` } : {}) },
  });
  if (response.status === 204) return null;
  const data = await response.json();
  if (!response.ok) throw new Error(data.error || "Не удалось выполнить запрос");
  return data;
}

export function roomSession(code) {
  try { return JSON.parse(sessionStorage.getItem(`room:${code.toUpperCase()}`)); }
  catch { return null; }
}

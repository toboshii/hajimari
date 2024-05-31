import { dev } from "$app/environment";

const base = dev ? "http://localhost:3000" : "";

export function api(
  fetch: Function,
  method: string,
  resource: string,
  data?: Record<string, unknown>,
) {
  return fetch(`${base}/${resource}`, {
    method,
    headers: {
      "content-type": "application/json",
    },
    body: data && JSON.stringify(data),
  });
}

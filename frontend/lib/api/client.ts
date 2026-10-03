export class ApiTransportError extends Error {
  constructor() {
    super("The request could not be completed.");
    this.name = "ApiTransportError";
  }
}

export async function apiRequest(
  path: string,
  init: RequestInit = {},
): Promise<Response> {
  try {
    const headers = new Headers(init.headers);
    headers.set("Accept", "application/json");
    return await fetch(path, {
      ...init,
      headers,
    });
  } catch {
    throw new ApiTransportError();
  }
}

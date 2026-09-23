export class ApiTransportError extends Error {
  constructor() {
    super("The request could not be completed.");
    this.name = "ApiTransportError";
  }
}

export async function apiRequest(path: string): Promise<Response> {
  try {
    return await fetch(path, {
      headers: {
        Accept: "application/json",
      },
    });
  } catch {
    throw new ApiTransportError();
  }
}

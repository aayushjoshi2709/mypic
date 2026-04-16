import status from "http-status";
import toast from "react-hot-toast";


export const UNAUTHORIZED_EVENT = "unauthorized";


class ApiClient {
  async handleStatus(response: Promise<Response>) {
    const res = await response;
    if (res.status === status.UNAUTHORIZED) {
      toast.error("Unauthorized. Please log in again.");
      window.dispatchEvent(new Event(UNAUTHORIZED_EVENT));
      return;
    }
    if (res.status === status.BAD_REQUEST) {
      const errorData = await res.json();
      toast.error(errorData.message);
    }

    if (!res.ok) {
      const errorData = await res.json();
      toast.error('An error occurred. Please try again.');
      throw new Error(
        `API request failed with status ${res.status}: ${errorData.message || res.statusText}`,
      );
    }
    return res.json();
  }

  async handleError(func: () => Promise<Response>) {
    try {
      return this.handleStatus(func());
    } catch (error) {
      console.error("API POST request failed:", error);
      toast.error("An error occurred. Please try again.");
      throw error;
    }
  }

  async post(url: string, data: unknown) {
    return this.handleError(async () => {
      return fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
    });
  }

  async get(url: string) {
    return this.handleError(async () => {
      return fetch(url);
    });
  }

  async put(url: string, data: unknown) {
    return this.handleError(async () => {
      return fetch(url, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
    });
  }

  async delete(url: string) {
    return this.handleError(async () => {
      return fetch(url, {
        method: "DELETE",
      });
    });
  }
}

export default ApiClient;

export const apiClientObj = Object.freeze(new ApiClient());

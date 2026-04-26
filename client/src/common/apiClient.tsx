import status from "http-status";
import toast from "react-hot-toast";


export const UNAUTHORIZED_EVENT = "unauthorized";


class ApiClient {
  async prepareHeaders() {
    const token = localStorage.getItem("token");
    return {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    };
  }
  async handleStatus(response: Promise<Response>) {
    const res = await response;
    if (res.status === status.UNAUTHORIZED) {
      toast.error("Unauthorized. Please log in again.");
      window.dispatchEvent(new Event(UNAUTHORIZED_EVENT));
      return;
    }
    if (res.status === status.BAD_REQUEST) {
      const errorData = await res.json();
      toast.error(errorData.error);
      throw new Error(
        `API request failed with status ${res.status}: ${errorData.error}`,
      );
    }

    if (res.status >= 400) {
      const errorData = await res.json();
      toast.error("An error occurred. Please try again.");
      throw new Error(
        `API request failed with status ${res.status}: ${errorData.message || res.statusText}`,
      );
    }

    if(res.status == 204){
      return;
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
        headers: await this.prepareHeaders(),
        body: JSON.stringify(data),
      });
    });
  }

  async get(url: string) {
    return this.handleError(async () => {
      return fetch(url, {
        headers: await this.prepareHeaders(),
      });
    });
  }

  async put(url: string, data: unknown) {
    return this.handleError(async () => {
      return fetch(url, {
        method: "PUT",
        headers: await this.prepareHeaders(),
        body: JSON.stringify(data),
      });
    });
  }

  async delete(url: string) {
    return this.handleError(async () => {
      return fetch(url, {
        method: "DELETE",
        headers: await this.prepareHeaders(),
      });
    });
  }
}

export default ApiClient;

export const apiClientObj = Object.freeze(new ApiClient());

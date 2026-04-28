import { useState } from "react";
import { apiClientObj } from "../common/apiClient";
import { routes } from "../common/routes";

const useImageDialog = () => {
  const [s3Key, setS3Key] = useState<string>("");
  const [originalFileName, setOriginalFileName] = useState<string>("");
  const [uploadInProgress, setUploadInProgress] = useState<boolean>(false);

  async function getPresignedUrl(name: string, type: string) {
    return await apiClientObj.post(routes.GET_PRESIGNED_URL, {
      originalName: name,
      type,
    });
  }

  async function uploadImageToS3(url: string, file: File) {
    const response = await fetch(url, {
      method: "PUT",
      headers: {
        "Content-Type": file.type,
      },
      body: file,
    });

    if (!response.ok) {
      throw new Error("Upload failed");
    }
  }

  async function onFileChange(event: Event) {
    setUploadInProgress(true);
    const input = event.target as HTMLInputElement | null;
    const file = input?.files?.[0];

    if (!file) return;

    try {
      const imageOrVideo = file.type.startsWith("image/")
        ? "images"
        : "videos";

      const { url, key } = await getPresignedUrl(
        file.name,
        imageOrVideo
      );

      await uploadImageToS3(url, file);

      setOriginalFileName(file.name);
      setS3Key(key);
      setUploadInProgress(false);
    } catch (error) {
      console.error(error);
      setUploadInProgress(false);
    }
  }

  function openFileDialog() {
    const fileInput = document.createElement("input");
    fileInput.type = "file";
    fileInput.accept = "image/*";
    fileInput.multiple = false;

    fileInput.onchange = onFileChange;

    fileInput.click();
  }

  return [openFileDialog, s3Key, originalFileName, uploadInProgress] as const;
};

export default useImageDialog;
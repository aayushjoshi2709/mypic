import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";

const Upload = () => {
  async function getPresinedUrl(fileName: string, fileType: string) {
    return await apiClientObj.post(routes.GET_PRESIGNED_URL, {
      originalName: fileName,
      type: fileType,
    });
  }

  async function uploadImageToS3(url: string, file: File) {
    await fetch(url, {
      method: "PUT",
      headers: {
        "Content-Type": file.type,
      },
      body: file,
    });
  }

  async function updateImageDatabase(url: string) {
    await apiClientObj.post(routes.CREATE_IMAGE, {
      url,
    });
  }

  async function openFileDialog() {
    const fileInput = document.createElement("input");
    fileInput.type = "file";
    fileInput.accept = "image/*";
    fileInput.multiple = false;
    fileInput.onchange = async (event) => {
      const files = (event.target as HTMLInputElement).files;
      const file = files ? files[0] : null;
      if (file) {
        const fileName = file.name;
        const imageOrVedio = file.type.startsWith("image/")
          ? "images"
          : "videos";
        const { url } = await getPresinedUrl(fileName, imageOrVedio);
        await uploadImageToS3(url, file);
        await updateImageDatabase(url);
      }
    };
    fileInput.click();
  }
  return (
    <div className="h-100 m-4 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
      <h1 className="text-4xl font-bold mb-4">Drag and drop to upload</h1>
      <p>Or click the button</p>
      <button
        onClick={openFileDialog}
        className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
      >
        Upload
      </button>
    </div>
  );
};

export default Upload;

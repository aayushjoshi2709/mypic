import { useDispatch } from "react-redux";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";
import toast from "react-hot-toast";
import { useEffect, useState } from "react";
import { setFetchImages } from "../../store/image.slice";
import useImageDialog from "../../customHooks/useImageDialog";

const Upload = () => {
  const dispatch = useDispatch();
  const [loading, setLoading] = useState(false);
  const [openFileDialog, s3Key, originalFileName, uploadInProgress] = useImageDialog();
  
  useEffect(()=>{
    if(uploadInProgress === false && s3Key && originalFileName){
      async function uploadFile(){
        await apiClientObj.post(routes.CREATE_IMAGE, {
          originalName: originalFileName,
          key: s3Key,
        });
        dispatch(setFetchImages());
        setLoading(false);
        toast.success("Image uploaded successfully!");
      }
      uploadFile();
    }
  },[dispatch, originalFileName, s3Key, uploadInProgress])



  async function uploadImage(){
    setLoading(true);
    openFileDialog();
  }
  return (
    <div className="h-100 m-4 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
      <h1 className="text-4xl font-bold mb-4">Drag and drop to upload</h1>
      <p>Or click the button</p>
      <button
        onClick={uploadImage}
        className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
        disabled={loading}
      >
        {loading ? "Uploading..." : "Upload"}
      </button>
    </div>
  );
};

export default Upload;

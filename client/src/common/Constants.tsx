export const ModalNames = {
  DELETE_MODAL: "DELETE_MODAL",
  PREVIEW_MODAL: "PREVIEW_MODAL",
  ADD_GROUP: "ADD_GROUP",
  ADD_IMAGE_TO_GROUP: "ADD_IMAGE_TO_GROUP",
} as const;

export const ImageLoadTypes = {
  IMAGE: "IMAGE",
  GROUP: "GROUP",
} as const;

export type ImageLoadType =
  (typeof ImageLoadTypes)[keyof typeof ImageLoadTypes];

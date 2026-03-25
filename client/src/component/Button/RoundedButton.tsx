interface ButtonProps {
  text: string;
}

export const RoundedButtonPrimary = (button: ButtonProps) => {
  return (
    <input
      className="p-2 px-8 hover:bg-green-600 bg-green-500 drop-shadow-2xl rounded-full text-white font-semibold"
      type="submit"
      value={button.text}
    />
  );
};

export const RoundedButtonSecondary = (button: ButtonProps) => {
  return (
    <input
      className="p-2 px-8 hover:bg-blue-600 bg-blue-500 drop-shadow-2xl rounded-full text-white font-semibold"
      type="button"
      value={button.text}
    />
  );
};

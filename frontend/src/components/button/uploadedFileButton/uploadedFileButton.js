export default function UploadedFileButton(state) {
  return (
    <>
      <button className={state.className} onClick={() => state.onClick()}>
        Uploaded File
      </button>
    </>
  );
}

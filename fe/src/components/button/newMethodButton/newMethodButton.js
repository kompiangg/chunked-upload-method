export default function NewMethodButton(state) {
  return (
    <>
      <button className={state.className} onClick={() => state.onClick()}>
        New Method
      </button>
    </>
  );
}

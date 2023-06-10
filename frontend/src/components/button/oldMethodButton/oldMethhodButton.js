export default function OldMethodButton(state) {
  return (
    <>
      <button className={state.className} onClick={() => state.onClick()}>
        Old Method
      </button>
    </>
  );
}

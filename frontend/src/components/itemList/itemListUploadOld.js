import './itemList.css';
import docs from '../../assets/docs.png';
import { fetch } from '../../utils/axios';

export default function ItemListUploadOld({ objOfArr }) {
  const cells = objOfArr.map((v) => {
    return (
      <div className="wrap-file-list" key={v.key}>
        <img src={docs} alt="Docs Icon" />

        <div className="wrap-file-metadata">
          <p>
            <strong>{v.fileName}</strong>
          </p>
          <p>File Size: {v.fileSize} Bytes</p>
        </div>

        <button
          className="menu-button-active"
          onClick={(e) => {
            handleOnClick(e, v.file);
          }}
        >
          Upload
        </button>
      </div>
    );
  });

  return <div>{cells}</div>;
}

async function handleOnClick(e, file) {
  console.log(e.target.className);

  e.target.classList.toggle('upload-button-uploading');

  e.target.innerHTML = 'Process';
  e.target.setAttribute('disabled', true);

  const body = new FormData();

  body.append('file', file);

  const res = await fetch.post('/v1/upload', body, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });

  if (!res.error) {
    e.target.innerHTML = 'Success';
  } else {
    e.target.innerHTML = 'Failed';
  }

  e.target.classList.toggle('upload-button-uploading');
  e.target.classList.add('menu-button-final-state');
}

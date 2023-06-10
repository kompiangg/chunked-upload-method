import { useEffect, useState } from 'react';
import './itemList.css';
import docs from '../../assets/docs.png';
import { fetch } from '../../utils/axios';

export default function ItemListUpload() {
  const [uploadedFile, setUploadedFile] = useState([]);

  useEffect(() => {
    (async function () {
      const data = await fetch.get('/v1/upload');
      setUploadedFile(data.data);
    })();
  }, []);

  const cells = uploadedFile.map((v) => {
    return (
      <div className="wrap-file-list" key={v.id}>
        <img src={docs} alt="Docs Icon" />

        <div className="wrap-file-metadata">
          <p>
            <strong>{v.name}</strong>
          </p>
          <p>Created at: {v.created_at}</p>
        </div>

        <a href={v.url} rel="noopener" target="_blank">
          open
        </a>
      </div>
    );
  });

  return <div>{cells}</div>;
}

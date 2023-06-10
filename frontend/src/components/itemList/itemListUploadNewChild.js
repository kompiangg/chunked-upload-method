import { useState } from 'react';
import docs from '../../assets/docs.png';
import { fetch } from '../../utils/axios';

export default function ItemListUploadNewChild({ file }) {
  const [percentage, setPercentage] = useState(0);

  return (
    <div className="wrap-file-list" key={file.idx}>
      <img src={docs} alt="Docs Icon" />

      <div className="wrap-file-metadata">
        <p>
          <strong>{file.fileName}</strong>
        </p>
        <p>File Size: {file.fileSize} Bytes</p>
      </div>

      <button
        onClick={(e) => handleOnClick(e, file.file, percentage, setPercentage)}
        className="menu-button-active"
      >
        {percentage === 0 ? 'Upload' : `${percentage}%`}
      </button>
    </div>
  );
}

async function handleOnClick(e, file, percentage, setPercentage) {
  e.target.classList.toggle('upload-button-uploading');
  e.target.setAttribute('disabled', true);

  const fileReader = new FileReader();

  try {
    fileReader.onload = async (e) => {
      await uploadChunkData(e, file, percentage, setPercentage);
    };

    fileReader.readAsArrayBuffer(file);
  } catch (error) {
    throw error;
  } finally {
    e.target.classList.toggle('upload-button-uploading');
    e.target.classList.add('menu-button-final-state');
  }
}

async function uploadChunkData(e, file, percentage, setPercentage) {
  try {
    const metadata = await getUploadMetadata(
      e.target.result.byteLength,
      file.name
    );

    for (let idx = 0; idx < metadata.chunk_count; idx++) {
      const binaryArr = e.target.result.slice(
        idx * metadata.chunk_byte_size,
        idx * metadata.chunk_byte_size + metadata.chunk_byte_size
      );

      await sendChunkData(metadata.unique_name, binaryArr, idx);

      setPercentage(Math.trunc((idx / metadata.chunk_count) * 100));
    }
    setPercentage(95);

    await finishSendChunkData(metadata.unique_name);
    setPercentage(100);
  } catch (error) {
    throw Error(error);
  }
}

async function getUploadMetadata(byteLength, fileName) {
  try {
    const res = await fetch.post(
      '/v2/upload/request',
      {
        size: byteLength,
        origin_name: fileName,
      },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      }
    );

    return res.data;
  } catch (error) {
    throw error;
  }
}

async function sendChunkData(fileName, chunkedData, chunkOrder) {
  try {
    const res = await fetch.post('/v2/upload', chunkedData, {
      headers: {
        'Content-Type': 'application/octet-stream',
        'X-Identity-Name': fileName,
        'X-Chunk-Order': chunkOrder,
      },
    });

    return res.data;
  } catch (error) {
    throw error;
  }
}

async function finishSendChunkData(uniqueName) {
  try {
    const res = await fetch.post(
      '/v2/upload/finish',
      { unique_name: uniqueName },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      }
    );

    return res.data;
  } catch (error) {
    throw error;
  }
}

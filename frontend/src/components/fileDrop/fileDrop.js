import React, { useCallback, useEffect, useReducer, useRef } from 'react';
import './fileDrop.css';

export default function FileDrop({ fileList, setFileList }) {
  const onDragOver = (event) => {
    event.preventDefault();
  };

  const ref = useRef(null);

  const onDragLeave = (event) => {
    ref.current.className = 'label-file-upload';
  };

  const onDragEnter = (event) => {
    event.preventDefault();
    ref.current.className = 'label-file-upload hover';
  };

  const onDrop = (event) => {
    event.preventDefault();
    ref.current.className = 'label-file-upload';

    const files = event.dataTransfer.files;
    handleFiles(files);
  };

  const onChange = (event) => {
    const files = event.target.files;
    handleFiles(files);
  };

  const handleFiles = useCallback((files) => {
    for (let i = 0; i < files.length; i++) {
      console.log(files[i]);
    }
    setFileList([...files]);
  }, []);

  return (
    <div className="file-upload">
      <div className="wrapper-file-upload">
        <label
          ref={ref}
          htmlFor="fileInput"
          className="label-file-upload"
          style={{ fontWeight: '600', width: '100%', height: '100%' }}
          onDragOver={onDragOver}
          onDrop={onDrop}
          onDragEnter={onDragEnter}
          onDragLeave={onDragLeave}
        >
          <span style={{ fontSize: '30px', fontWeight: 600 }}>+</span> Drag and
          drop your file here or click to upload
        </label>
        <input id="fileInput" type="file" onChange={onChange} />
      </div>
    </div>
  );
}

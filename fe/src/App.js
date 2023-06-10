import { useState } from 'react';
import './App.css';
import Footer from './components/footer/footer';
import Header from './components/header/header';
import Menu from './components/menu/menu';
import FileDrop from './components/fileDrop/fileDrop';
import ItemListUploaded from './components/itemList/ItemListUploaded';
import ItemListUploadOld from './components/itemList/itemListUploadOld';
import ItemListUploadNew from './components/itemList/itemListUploadNew';

export default function App() {
  const [activeButton, setActiveButton] = useState(1);
  const [fileList, setFileList] = useState([]);

  console.log(process.env.REACT_APP_BASE_API_URL);

  return (
    <div className="App">
      <div className="body">
        <Header></Header>
        <Menu
          activeButton={activeButton}
          setActiveButton={setActiveButton}
          setFileList={setFileList}
        ></Menu>

        {content(fileList, setFileList, activeButton, setActiveButton)}
      </div>
      <Footer></Footer>
    </div>
  );
}

function content(fileList, setFileList, activeButton, setActiveButton) {
  if (fileList.length === 0 && (activeButton === 1 || activeButton === 2)) {
    return (
      <FileDrop
        activeButton={activeButton}
        setActiveButton={setActiveButton}
        fileList={fileList}
        setFileList={setFileList}
      ></FileDrop>
    );
  } else if (fileList.length !== 0 && activeButton === 1) {
    const objOfArr = fileList.map((v, idx) => {
      return {
        fileName: v.name,
        fileSize: v.size,
        idx: idx,
        file: v,
      };
    });

    return (
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          marginBottom: '30px',
        }}
      >
        <ItemListUploadOld objOfArr={objOfArr}></ItemListUploadOld>
      </div>
    );
  } else if (fileList.length !== 0 && activeButton === 2) {
    const objOfArr = fileList.map((v, idx) => {
      return {
        fileName: v.name,
        fileSize: v.size,
        idx: idx,
        file: v,
      };
    });

    return (
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          marginBottom: '30px',
        }}
      >
        <ItemListUploadNew objOfArr={objOfArr}></ItemListUploadNew>
      </div>
    );
  }

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        marginBottom: '30px',
      }}
    >
      <ItemListUploaded></ItemListUploaded>
    </div>
  );
}

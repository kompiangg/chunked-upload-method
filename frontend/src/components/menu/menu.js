import './menu.css';
import OldMethodButton from '../button/oldMethodButton/oldMethhodButton';
import { menu } from '../../constant/menu';
import NewMethodButton from '../button/newMethodButton/newMethodButton';
import UploadedFileButton from '../button/uploadedFileButton/uploadedFileButton';

export default function Menu({ activeButton, setActiveButton, setFileList }) {
  let activeButtonTag = 'menu-button-active';

  const handleOnClick = (buttonId) => {
    setActiveButton(buttonId);
    setFileList([]);
  };

  return (
    <div className="menu">
      <OldMethodButton
        className={`${
          activeButton === menu.oldMethodButton
            ? activeButtonTag
            : 'menu-button'
        }`}
        onClick={() => handleOnClick(menu.oldMethodButton)}
      ></OldMethodButton>
      <NewMethodButton
        className={`${
          activeButton === menu.newMethodButton
            ? activeButtonTag
            : 'menu-button'
        }`}
        onClick={() => handleOnClick(menu.newMethodButton)}
      ></NewMethodButton>
      <UploadedFileButton
        className={`${
          activeButton === menu.uploadedFileButton
            ? activeButtonTag
            : 'menu-button'
        }`}
        onClick={() => handleOnClick(menu.uploadedFileButton)}
      ></UploadedFileButton>
    </div>
  );
}

import './itemList.css';
import ItemListUploadNewChild from './itemListUploadNewChild';

export default function ItemListUploadNew({ objOfArr }) {
  return (
    <div>
      {objOfArr.map((v) => (
        <ItemListUploadNewChild file={v} key={v.idx}></ItemListUploadNewChild>
      ))}
    </div>
  );
}

import shipperLogo from '../../assets/shipper-logo.webp';
import './header.css';

export default function Header() {
  return (
    <div className="header">
      <img src={shipperLogo} alt="Shipper Logo" />

      <h1 style={{ marginBottom: '0px' }}>Demo Final Project</h1>
      <h2>
        Implementation Percent-Done Progress Indicator with Cached-Chunked Data
        Transfer Method
      </h2>
    </div>
  );
}

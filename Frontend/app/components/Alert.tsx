'use client';
import { useEffect } from 'react';

export type ALERT_TYPE = 'success' | 'error' | 'info';
export type ALERT_PROPS = {
  message?: string;
  type: ALERT_TYPE;
  onClose: () => void;
};

const alertStyles = {
  success: 'bg-green-100 text-green-800 border-green-300',
  error: 'bg-red-100 text-red-800 border-red-300',
  info: 'bg-blue-100 text-blue-800 border-blue-300',
};

const Alert: React.VFC<ALERT_PROPS> = ({ message, type = 'info', onClose }) => {
  useEffect(() => {
    if (!message) return;
    const timer = setTimeout(onClose, 3000);

    return () => clearTimeout(timer);
  }, [message, onClose]);
  if (!message) return null;
  return (
    <div className="fixed top-2 w-screen flex justify-center" role="alert">
      <div className={`px-4 py-2 mb-4 z-20 rounded ${alertStyles[type]}`}>
        <div className="flex justify-between items-center">
          <span>{message}</span>
        </div>
      </div>
    </div>
  );
};

export default Alert;

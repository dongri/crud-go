// app/LoadingBar.tsx
'use client';
import { useContext } from 'react';
import { LoadingContext } from './loading-context';

export default function LoadingBar() {
  const { loading } = useContext(LoadingContext);

  return (
    <>
      {loading && (
        <div style={{
          position: 'fixed',
          top: 0,
          left: 0,
          width: '100%',
          height: '4px',
          backgroundColor: '#0070f3',
          zIndex: 9999,
          animation: 'loadingAnimation 2s linear infinite'
        }} />
      )}
      <style jsx>{`
        @keyframes loadingAnimation {
          0% { transform: translateX(-100%); }
          100% { transform: translateX(100%); }
        }
      `}</style>
    </>
  );
}

// app/loading-context.tsx
'use client';
import React, { createContext, useState, ReactNode } from 'react';

interface LoadingContextType {
  loading: boolean;
  setLoading: (state: boolean) => void;
}

export const LoadingContext = createContext<LoadingContextType>({
  loading: false,
  setLoading: () => {},
});

export function LoadingProvider({ children }: { children: ReactNode }) {
  const [loading, setLoading] = useState(false);
  return (
    <LoadingContext.Provider value={{ loading, setLoading }}>
      {children}
    </LoadingContext.Provider>
  );
}

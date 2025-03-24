'use client';

import { useContext } from 'react';
import { LoadingContext } from '../app/loading-context';

const BASE_URL = process.env.NEXT_APP_BASE_URL || 'http://localhost:3001/api';

//======================= Fetch with XXXXX =======================//
const fetchWithPublic = async (path: string, options: RequestInit = {}) => {
  
  const headers = {
    'Content-Type': 'application/json',
    ...options.headers,
  };

  try {
    const response = await fetch(`${BASE_URL}${path}`, {
      ...options,
      headers,
    });
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(
        errorData.message || `HTTP error! status: ${response.status}`
      );
    }
    return await response.json();
  } catch (error) {
    console.error(`Fetch error on ${path}:`, error);
    return { error: error };
  }
};

//======================= Http methods with XXXX =======================//
export function useClient() {
  const { setLoading } = useContext(LoadingContext);

  const publicGet = async (path: string) => {    
    setLoading(true);
    try {
      // sleep for 1 second
      // await new Promise(resolve => setTimeout(resolve, 1000));
      return fetchWithPublic(path, {
        method: 'GET',
      });
    } catch (error) {
      console.error('PublicGet error:', error);
      return { error: error };
    } finally {
      setLoading(false);
    }
  };

  const publicPost = async (path: string, body: Record<string, unknown>) => {
    setLoading(true);
    try {
      return fetchWithPublic(path, {
        method: 'POST',
        body: JSON.stringify(body),
      });
    } catch (error) {
      console.error('PublicPost error:', error);
      return { error: error };
    } finally {
      setLoading(false);
    }
  };

  return { publicGet, publicPost };
}

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
export const PublicGet = (path: string) => {
  return fetchWithPublic(path, {
    method: 'GET',
  });
};

export const PublicPost = (path: string, body: Record<string, unknown>) => {
  return fetchWithPublic(path, {
    method: 'POST',
    body: JSON.stringify(body),
  });
};

const { VITE_API_ENDPOINT: endpoint } = process.env;

const api = async (url, method, body = {}) => {
  const { fetch } = window;

  const headers = {
    method,
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
    },
  };

  const updateMethods = ['PUT', 'PATCH'];
  if (updateMethods.includes(method)) {
    headers.body = JSON.stringify(body);
  }

  return await fetch(`${endpoint}/${url}`, headers);
};

export default api;

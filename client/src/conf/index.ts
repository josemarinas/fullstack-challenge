interface ConfObj {
  baseUrl: string,
  ipfsUrl: string
}

export const Conf:ConfObj = {
  baseUrl: process.env.REACT_APP_API_URL || 'http://localhost:8000/api',
  ipfsUrl: process.env.REACT_APP_IPFS_URL|| '/ip4/0.0.0.0/tcp/5001'
}
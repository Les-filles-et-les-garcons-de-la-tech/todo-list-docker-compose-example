export const environment = {
  production: true,
  apiUrl: (window as { [key: string]: any })["env"]["backUrl"] as string || "http://localhost:80"
};

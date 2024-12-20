import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

export default defineConfig({
  plugins: [react()],
  root: path.resolve(__dirname, '../'), // Actualiza esta línea
  resolve: {
    alias: {
      '@': path.resolve(__dirname, '../src')
    }
  },
  build: {
    outDir: path.resolve(__dirname, '../dist'), // Asegúrate de que la salida esté en el directorio correcto
  }
});
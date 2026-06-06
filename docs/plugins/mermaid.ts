import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));

export function pluginMermaid() {
  return {
    name: 'plugin-mermaid',
    globalUIComponents: [join(__dirname, '../components/Mermaid.tsx')],
  };
}

import { useEffect } from 'react';

declare global {
  interface Window {
    mermaid: any;
  }
}

export default function Mermaid() {
  useEffect(() => {
    if (typeof window === 'undefined') return;

    const MERMAID_CDN = 'https://cdn.jsdelivr.net/npm/mermaid@11/dist/mermaid.min.js';

    const renderDiagrams = () => {
      const m = window.mermaid;
      if (!m) return;

      document.querySelectorAll('pre code.language-mermaid').forEach((block) => {
        const pre = block.parentElement;
        if (!pre || pre.getAttribute('data-mermaid')) return;
        pre.setAttribute('data-mermaid', '1');
        try {
          const id = `m${Date.now()}${Math.random().toString(36).slice(2, 6)}`;
          m.render(id, (block.textContent || '').trim()).then(({ svg }: { svg: string }) => {
            const div = document.createElement('div');
            div.style.textAlign = 'center';
            div.style.margin = '1.5rem 0';
            div.style.overflow = 'auto';
            div.innerHTML = svg;
            pre.replaceWith(div);
          });
        } catch {
          /* skip broken diagrams */
        }
      });
    };

    if (!window.mermaid) {
      const script = document.createElement('script');
      script.src = MERMAID_CDN;
      script.async = true;
      script.onload = () => {
        window.mermaid.initialize({
          startOnLoad: false,
          theme: 'default',
          securityLevel: 'loose',
        });
        renderDiagrams();
      };
      document.head.appendChild(script);
    } else {
      renderDiagrams();
    }

    // Re-render on client navigation (Rspress SPA transitions)
    const mo = new MutationObserver(() => {
      setTimeout(renderDiagrams, 400);
    });
    mo.observe(document.body, { childList: true, subtree: true });

    return () => mo.disconnect();
  }, []);

  return null;
}

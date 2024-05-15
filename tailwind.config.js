/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      'internal/views/**/*.templ',
    ],
    darkMode: 'class',
    theme: {
      extend: {
        fontFamily: {
          mono: ['Courier Prime', 'monospace'],
        }
      },
    },
    plugins: [
      require('@tailwindcss/forms'),
    ],
    corePlugins: {
      preflight: true,
    }
  }
  
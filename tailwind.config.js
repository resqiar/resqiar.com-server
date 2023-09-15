/** @type {import('tailwindcss').Config} */
export default {
  content: ["./views/**/*.html"],
  theme: {
    extend: {
      colors: {
        brand: {
          primary: '#995600',
          light: '#ffa500'
        }
      }
    },
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ['lofi'],
  }
}


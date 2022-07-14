/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/typography')],
  purge: ['./src/components/**/*.tsx','./src/pages/**/*.tsx'],
	darkMode: "class",
	mode: "jit"
}

module.exports = {
  content: ['./app/**/*.{hbs,js}'],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["garden", "forest"],
  },
}

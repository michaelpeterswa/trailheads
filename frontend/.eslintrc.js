module.exports = {
    env: {
        "browser": true,
        "node": true,
        "es2021": true
    },
    parser: "babel-eslint",
    plugins: ['ember'],
    extends: [
      'eslint:recommended',
      'plugin:ember/recommended' // or other configuration
    ],
    ignorePatterns: ["dist/**"],
    rules: {
    }
}

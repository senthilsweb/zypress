/*
** TailwindCSS Configuration File
**
** Docs: https://tailwindcss.com/docs/configuration
** Default: https://github.com/tailwindcss/tailwindcss/blob/master/stubs/defaultConfig.stub.js
*/
const defaultTheme = require('tailwindcss/defaultTheme')
const colors = require('tailwindcss/colors')
const plugin = require('tailwindcss/plugin')
module.exports = {
  important: '#app',
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        black: colors.black,
        white: colors.white,
        gray: colors.gray,
        red: colors.red,
        yellow: colors.yellow,
        green: colors.teal,
        cyan:colors.cyan,
        blue: colors.blue,
        indigo: colors.indigo,
        purple: colors.purple,
        pink: colors.pink,
        rose: colors.rose,
        teal: colors.teal,
        orange: colors.orange,
        magenta:colors.magenta
      },
      spacing: {
        '80': '20rem',
        '108': '27rem',
        '160': '40rem',
      },
      borderWidth: {
        '14': '14px',
      },
      container: {
        padding: '1rem'
      },
      fontFamily: {
        sans: [
          'Inter var',
          'Nunito Sans',
          'Roboto',
          '-apple-system',
          'BlinkMacSystemFont',
          '"Segoe UI"',
          '"Helvetica Neue"',
          'Arial',
          '"Noto Sans"',
          'sans-serif',
          '"Apple Color Emoji"',
          '"Segoe UI Emoji"',
          '"Segoe UI Symbol"',
          '"Noto Color Emoji"',
        ]
      },
      inset: {
        '-16': '-4rem',
      },
      fontSize: {
        xxs: '0.625rem'
      },
      boxShadow: {
        'outline-green': '0 0 0 3px rgba(5,202,160, 0.8)',
      },
    }
  },
  variants: {
    textColor: ['responsive', 'hover', 'focus', 'group-hover'],
    backgroundColor: ['responsive', 'hover', 'focus', 'even', 'odd'],
  },
  plugins: [
    require('@tailwindcss/ui')({
      layout: 'sidebar',
    }),
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    plugin(function({ addComponents, theme }) {
      const buttons = {
        '.page': {
          backgroundColor: theme('colors.green.500'),
          borderColor: theme('colors.green.500'),
          color: theme('colors.white'),
          '&:hover': {
            backgroundColor: theme('colors.green.600'),
            borderColor: theme('colors.green.600')
          },
        },
      }
      addComponents(buttons)
    })
  ],future: {
    removeDeprecatedGapUtilities: true,
  },
  purge: {
    enabled: process.env.NODE_ENV === 'production',
    content: [
      'components/**/*.vue',
      'layouts/**/*.vue',
      'pages/**/*.vue',
      'plugins/**/*.js',
      'nuxt.config.js',
      // TypeScript
      'plugins/**/*.ts',
      'nuxt.config.ts'
    ]
  }
}

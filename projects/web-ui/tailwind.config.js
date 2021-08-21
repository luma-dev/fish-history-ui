import colors from 'windicss/colors';
import typography from 'windicss/plugin/typography';

export default {
  darkMode: 'class',
  plugins: [typography],
  theme: {
    extend: {
      colors: {
        teal: colors.teal,
        svelte: {
          500: '#ff3e00',
        },
      },
    },
  },
};

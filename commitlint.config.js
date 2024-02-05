module.exports = {
  parserPreset: {
    parserOpts: {
      headerPattern: /^\[(\w+)\]\[(\d+)\]: (.+)/,
      headerCorrespondence: ['type', 'scope', 'subject'],
    },
  },
  plugins: [
    {
      rules: {
        'header-match-team-pattern': parsed => {
          const {type, scope, subject} = parsed;
          if (type === null && scope === null && subject === null) {
            return [
              false,
              'header must be in format [tag][task number]: commit message, tag must be fix, implement, enhance',
            ];
          }
          return [true, ''];
        },
      },
    },
  ],
  rules: {
    'header-match-team-pattern': [2, 'always'],
    'type-enum': [2, 'always', ['fix', 'implement', 'enhance']],
  },
};

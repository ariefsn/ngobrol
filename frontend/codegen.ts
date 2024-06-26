import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: 'http://localhost:6001/query',
  documents: ['src/**/*.gql.{ts,tsx}'],
  generates: {
    'src/graphql.generated.ts': {
      plugins: ['typescript', 'typescript-operations'],
    },
  },

}

export default config
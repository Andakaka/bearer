critical:
    - rule_dsrid: DSR-1
      rule_display_id: ruby_third_parties_sentry
      rule_description: Do not send sensitive data to Sentry.
      rule_documentation_url: https://curio.sh/reference/rules/ruby_third_parties_sentry
      line_number: 1
      filename: pkg/commands/process/settings/rules/ruby/third_parties/sentry/testdata/datatype_in_capture_message.rb
      category_groups:
        - PII
      parent_line_number: 1
      parent_content: 'Sentry.capture_message("test: #{user.email}")'
    - rule_dsrid: DSR-1
      rule_display_id: ruby_third_parties_sentry
      rule_description: Do not send sensitive data to Sentry.
      rule_documentation_url: https://curio.sh/reference/rules/ruby_third_parties_sentry
      line_number: 2
      filename: pkg/commands/process/settings/rules/ruby/third_parties/sentry/testdata/datatype_in_capture_message.rb
      category_groups:
        - PII
      parent_line_number: 2
      parent_content: 'Sentry.capture_message("test", extra: { email: user.email })'
    - rule_dsrid: DSR-1
      rule_display_id: ruby_third_parties_sentry
      rule_description: Do not send sensitive data to Sentry.
      rule_documentation_url: https://curio.sh/reference/rules/ruby_third_parties_sentry
      line_number: 3
      filename: pkg/commands/process/settings/rules/ruby/third_parties/sentry/testdata/datatype_in_capture_message.rb
      category_groups:
        - PII
      parent_line_number: 3
      parent_content: 'Sentry.capture_message("test", tags: { email: user.email })'
    - rule_dsrid: DSR-1
      rule_display_id: ruby_third_parties_sentry
      rule_description: Do not send sensitive data to Sentry.
      rule_documentation_url: https://curio.sh/reference/rules/ruby_third_parties_sentry
      line_number: 4
      filename: pkg/commands/process/settings/rules/ruby/third_parties/sentry/testdata/datatype_in_capture_message.rb
      category_groups:
        - PII
      parent_line_number: 4
      parent_content: 'Sentry.capture_message("test", user: { email: user.email })'


--

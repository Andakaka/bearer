data_types:
    - name: Email Address
      detectors:
        - name: javascript
          locations:
            - filename: pkg/commands/process/settings/rules/javascript/lang/exception/testdata/promise_reject.js
              line_number: 5
              field_name: email
              object_name: user
              subject_name: User
risks:
    - detector_id: javascript_lang_exception
      data_types:
        - name: Email Address
          stored: false
          locations:
            - filename: pkg/commands/process/settings/rules/javascript/lang/exception/testdata/promise_reject.js
              line_number: 5
              parent:
                line_number: 5
                content: Promise.reject(new PermissionDenied("fail" + user.email))
              field_name: email
              object_name: user
              subject_name: User
components: []


--

# Inline Form Idea

A sample and crude polyfill for the idea of inline HTML forms.

The idea is that there would be a new target type for HTML Forms, an attribute
of `target=_inline` would mean the that Form would be processed by the server
and the contents of the Form would be replaced with the HTML returned by the
server. That is, on submit the form values would be sent to the server, either
POST or GET, and the response should be HTML that the browser will simply
.innerHTML on the form that was submitted.

A [running example is available]( https://inline-form-nuau7zlm6q-uc.a.run.app). View
source on the main page to see a crude polyfill and an example form that uses
`target=_inline`.

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)
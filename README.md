make sure to run `pnpm install` and then `pnpm build` to get the dependencies. I will add a build to do the typescript and tailwindcss compiling eventaully.

This is the code for a rough website for my Dad's insurance brokering agency.

The CSS still needs to be worked on and pages would need to be added for processing different forms.

Otherwise it was cool try implement HTMX and have it hook into a GO program that processed the form and reactively altered the HTML.

It was also cool to set up the devcontainer to make an environment for typescript, go, and use the Caddy webserver.

The website is able to run as is using any webserver you need.

You just need to make sure to compile the goWebmailer that's included to run the form email sending.

And you need to make sure that you add in the webserver to point `/send` to the goWebmailer so that it can respond to the form requests.

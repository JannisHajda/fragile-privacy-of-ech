# doech
A Firefox extension for visualizing DoH and ECH usage on a per request basis.

## Build Instructions
### Requirements
* Node.js (v18+ recommended)
* npm

### Build steps
1. Open a terminal in the root directory of this source code folder.
2. Run `npm ci` to install the exact dependencies specified in the `package-lock.json` file.
3. Run the following command to compile and minify the CSS for production:

```bash
   npx @tailwindcss/cli -i ./src/styles/input.css -o ./src/styles/output.css --minify
```

## Third-Party Libraries
We rely on [jQuery](https://jquery.com/) and [Chart.js](https://www.chartjs.org/) for updating and visualization in the sidebar. These files (`src/lib/jquery.min.js`, `src/lib/chart.umd.min.js` and `src/lib/chart.umd.min.js.map`) are unmodified and were downloaded directly from their respective CDNs.
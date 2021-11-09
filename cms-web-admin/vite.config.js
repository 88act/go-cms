/* eslint-disable */
import legacyPlugin from '@vitejs/plugin-legacy';
// import usePluginImport from 'vite-plugin-importer';
import * as path from 'path';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import vuePlugin from '@vitejs/plugin-vue';
// @see https://cn.vitejs.dev/config/
export default ({
  command,
  mode
}) => {
  let NODE_ENV = process.env.NODE_ENV || 'development'
  let envFiles=[
    `.env.${NODE_ENV}`
  ]
  for (const file of envFiles) {
		const envConfig = dotenv.parse(fs.readFileSync(file))
		for (const k in envConfig) {
			process.env[k] = envConfig[k]
		}
	}

  let rollupOptions = {};


  let optimizeDeps = {};


  let alias = {
    '@': path.resolve(__dirname, './src'),
    'vue$': 'vue/dist/vue.runtime.esm-bundler.js',
  }

  let esbuild = {}

  return {
    base: './', // index.html文件所在位置
    root: './', // js导入的资源路径，src
    resolve: {
      alias,
    },
    define: {
      'process.env': {}
    },
    server: {
      open:true,
      port: process.env.VITE_CLI_PORT,
      proxy:{
       // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      [process.env.VITE_BASE_API]: { // 需要代理的路径   例如 '/api'
        target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
        changeOrigin: true,
        rewrite: path => path.replace(new RegExp('^' + process.env.VITE_BASE_API), ''),
      } 
      },
    },
    build: {
      target: 'es2015',
      minify: 'terser', // 是否进行压缩,boolean | 'terser' | 'esbuild',默认使用terser
      manifest: false, // 是否产出maifest.json
      sourcemap: false, // 是否产出soucemap.json
      outDir: 'dist', // 产出目录
      rollupOptions,
    },
    esbuild,
    optimizeDeps,
    plugins: [
      legacyPlugin({
        targets: ['Android > 39', 'Chrome >= 60', 'Safari >= 10.1', 'iOS >= 10.3', 'Firefox >= 54', 'Edge >= 15'],
      }), vuePlugin(),
    ],
    css: {
      preprocessorOptions: {
        less: {
          // 支持内联 JavaScript
          javascriptEnabled: true,
        }
      }
    },
  }
}
webpackHotUpdate_N_E("pages/torrent",{

/***/ "./pages/torrent.tsx":
/*!***************************!*\
  !*** ./pages/torrent.tsx ***!
  \***************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* WEBPACK VAR INJECTION */(function(module) {/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "default", function() { return torrent; });
/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ "./node_modules/react/jsx-dev-runtime.js");
/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var antd__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! antd */ "./node_modules/antd/es/index.js");
/* harmony import */ var _Component_filelist__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../Component/filelist */ "./Component/filelist.tsx");
/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/router */ "./node_modules/next/dist/client/router.js");
/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_3__);


var _jsxFileName = "/home/azama/ohkaca-react/pages/torrent.tsx",
    _s = $RefreshSig$();




function torrent() {
  return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Layout"], {
    children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Layout"].Content, {
      children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Row"], {
        align: "middle",
        justify: "center",
        style: {
          minHeight: "100vh"
        },
        children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Col"], {
          span: 18,
          children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Card"], {
            children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(FilelistWarpper, {}, void 0, false, {
              fileName: _jsxFileName,
              lineNumber: 12,
              columnNumber: 15
            }, this)
          }, void 0, false, {
            fileName: _jsxFileName,
            lineNumber: 11,
            columnNumber: 13
          }, this)
        }, void 0, false, {
          fileName: _jsxFileName,
          lineNumber: 10,
          columnNumber: 11
        }, this)
      }, void 0, false, {
        fileName: _jsxFileName,
        lineNumber: 9,
        columnNumber: 9
      }, this)
    }, void 0, false, {
      fileName: _jsxFileName,
      lineNumber: 8,
      columnNumber: 7
    }, this)
  }, void 0, false, {
    fileName: _jsxFileName,
    lineNumber: 7,
    columnNumber: 5
  }, this);
}

function FilelistWarpper() {
  _s();

  var hash = Object(next_router__WEBPACK_IMPORTED_MODULE_3__["useRouter"])().query.hash;

  if (!hash) {
    return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Skeleton"], {}, void 0, false, {
      fileName: _jsxFileName,
      lineNumber: 25,
      columnNumber: 12
    }, this);
  }

  console.log(hash);
  return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(_Component_filelist__WEBPACK_IMPORTED_MODULE_2__["default"], {
    hash: String(hash)
  }, void 0, false, {
    fileName: _jsxFileName,
    lineNumber: 29,
    columnNumber: 10
  }, this);
}

_s(FilelistWarpper, "CeygcqajjFExIxFEzW4x/gfWEGQ=", false, function () {
  return [next_router__WEBPACK_IMPORTED_MODULE_3__["useRouter"]];
});

_c = FilelistWarpper;

var _c;

$RefreshReg$(_c, "FilelistWarpper");

;
    var _a, _b;
    // Legacy CSS implementations will `eval` browser code in a Node.js context
    // to extract CSS. For backwards compatibility, we need to check we're in a
    // browser context before continuing.
    if (typeof self !== 'undefined' &&
        // AMP / No-JS mode does not inject these helpers:
        '$RefreshHelpers$' in self) {
        var currentExports = module.__proto__.exports;
        var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;
        // This cannot happen in MainTemplate because the exports mismatch between
        // templating and execution.
        self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.i);
        // A module can be accepted automatically based on its exports, e.g. when
        // it is a Refresh Boundary.
        if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {
            // Save the previous exports on update so we can compare the boundary
            // signatures.
            module.hot.dispose(function (data) {
                data.prevExports = currentExports;
            });
            // Unconditionally accept an update to this module, we'll check if it's
            // still a Refresh Boundary later.
            module.hot.accept();
            // This field is set when the previous version of this module was a
            // Refresh Boundary, letting us know we need to check for invalidation or
            // enqueue an update.
            if (prevExports !== null) {
                // A boundary can become ineligible if its exports are incompatible
                // with the previous exports.
                //
                // For example, if you add/remove/change exports, we'll want to
                // re-execute the importing modules, and force those components to
                // re-render. Similarly, if you convert a class component to a
                // function, we want to invalidate the boundary.
                if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {
                    module.hot.invalidate();
                }
                else {
                    self.$RefreshHelpers$.scheduleUpdate();
                }
            }
        }
        else {
            // Since we just executed the code for the module, it's possible that the
            // new exports made it ineligible for being a boundary.
            // We only care about the case when we were _previously_ a boundary,
            // because we already accepted this update (accidental side effect).
            var isNoLongerABoundary = prevExports !== null;
            if (isNoLongerABoundary) {
                module.hot.invalidate();
            }
        }
    }

/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(/*! ./../node_modules/webpack/buildin/harmony-module.js */ "./node_modules/webpack/buildin/harmony-module.js")(module)))

/***/ })

})
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vcGFnZXMvdG9ycmVudC50c3giXSwibmFtZXMiOlsidG9ycmVudCIsIm1pbkhlaWdodCIsIkZpbGVsaXN0V2FycHBlciIsImhhc2giLCJ1c2VSb3V0ZXIiLCJxdWVyeSIsImNvbnNvbGUiLCJsb2ciLCJTdHJpbmciXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUE7QUFDQTtBQUNBO0FBRWUsU0FBU0EsT0FBVCxHQUFtQjtBQUNoQyxzQkFDRSxxRUFBQywyQ0FBRDtBQUFBLDJCQUNFLHFFQUFDLDJDQUFELENBQVEsT0FBUjtBQUFBLDZCQUNFLHFFQUFDLHdDQUFEO0FBQUssYUFBSyxFQUFDLFFBQVg7QUFBb0IsZUFBTyxFQUFDLFFBQTVCO0FBQXFDLGFBQUssRUFBRTtBQUFFQyxtQkFBUyxFQUFFO0FBQWIsU0FBNUM7QUFBQSwrQkFDRSxxRUFBQyx3Q0FBRDtBQUFLLGNBQUksRUFBRSxFQUFYO0FBQUEsaUNBQ0UscUVBQUMseUNBQUQ7QUFBQSxtQ0FDRSxxRUFBQyxlQUFEO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFERjtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBREY7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQURGO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFERjtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBREY7QUFBQTtBQUFBO0FBQUE7QUFBQSxVQURGO0FBYUQ7O0FBRUQsU0FBU0MsZUFBVCxHQUEyQjtBQUFBOztBQUN6QixNQUFNQyxJQUFJLEdBQUdDLDZEQUFTLEdBQUdDLEtBQVosQ0FBa0JGLElBQS9COztBQUVBLE1BQUksQ0FBQ0EsSUFBTCxFQUFXO0FBQ1Qsd0JBQU8scUVBQUMsNkNBQUQ7QUFBQTtBQUFBO0FBQUE7QUFBQSxZQUFQO0FBQ0Q7O0FBRURHLFNBQU8sQ0FBQ0MsR0FBUixDQUFZSixJQUFaO0FBQ0Esc0JBQU8scUVBQUMsMkRBQUQ7QUFBVSxRQUFJLEVBQUVLLE1BQU0sQ0FBQ0wsSUFBRDtBQUF0QjtBQUFBO0FBQUE7QUFBQTtBQUFBLFVBQVA7QUFDRDs7R0FUUUQsZTtVQUNNRSxxRDs7O0tBRE5GLGUiLCJmaWxlIjoic3RhdGljL3dlYnBhY2svcGFnZXMvdG9ycmVudC5jM2NmMWE2OWY4ZjhiMGM4MTdiYi5ob3QtdXBkYXRlLmpzIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgTGF5b3V0LCBSb3csIENvbCwgQ2FyZCwgU2tlbGV0b24gfSBmcm9tIFwiYW50ZFwiO1xuaW1wb3J0IEZpbGVsaXN0IGZyb20gXCIuLi9Db21wb25lbnQvZmlsZWxpc3RcIjtcbmltcG9ydCB7IHVzZVJvdXRlciB9IGZyb20gXCJuZXh0L3JvdXRlclwiO1xuXG5leHBvcnQgZGVmYXVsdCBmdW5jdGlvbiB0b3JyZW50KCkge1xuICByZXR1cm4gKFxuICAgIDxMYXlvdXQ+XG4gICAgICA8TGF5b3V0LkNvbnRlbnQ+XG4gICAgICAgIDxSb3cgYWxpZ249XCJtaWRkbGVcIiBqdXN0aWZ5PVwiY2VudGVyXCIgc3R5bGU9e3sgbWluSGVpZ2h0OiBcIjEwMHZoXCIgfX0+XG4gICAgICAgICAgPENvbCBzcGFuPXsxOH0+XG4gICAgICAgICAgICA8Q2FyZD5cbiAgICAgICAgICAgICAgPEZpbGVsaXN0V2FycHBlciAvPlxuICAgICAgICAgICAgPC9DYXJkPlxuICAgICAgICAgIDwvQ29sPlxuICAgICAgICA8L1Jvdz5cbiAgICAgIDwvTGF5b3V0LkNvbnRlbnQ+XG4gICAgPC9MYXlvdXQ+XG4gICk7XG59XG5cbmZ1bmN0aW9uIEZpbGVsaXN0V2FycHBlcigpIHtcbiAgY29uc3QgaGFzaCA9IHVzZVJvdXRlcigpLnF1ZXJ5Lmhhc2g7XG5cbiAgaWYgKCFoYXNoKSB7XG4gICAgcmV0dXJuIDxTa2VsZXRvbiAvPjtcbiAgfVxuXG4gIGNvbnNvbGUubG9nKGhhc2gpO1xuICByZXR1cm4gPEZpbGVsaXN0IGhhc2g9e1N0cmluZyhoYXNoKX0gLz47XG59XG4iXSwic291cmNlUm9vdCI6IiJ9
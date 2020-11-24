webpackHotUpdate_N_E("pages/torrent",{

/***/ "./Component/filelist.tsx":
/*!********************************!*\
  !*** ./Component/filelist.tsx ***!
  \********************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* WEBPACK VAR INJECTION */(function(module) {/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "default", function() { return Filelist; });
/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ "./node_modules/react/jsx-dev-runtime.js");
/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var antd__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! antd */ "./node_modules/antd/es/index.js");
/* harmony import */ var swr__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! swr */ "./node_modules/swr/esm/index.js");
/* harmony import */ var _lib_graphql__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../lib/graphql */ "./lib/graphql.tsx");

var _jsxFileName = "/home/azama/ohkaca-react/Component/filelist.tsx";



function Filelist(props) {
  var _this = this;

  console.log(props);
  var columns = [{
    title: "Filename",
    key: "filename",
    dataIndex: "filename"
  }, {
    title: "Action",
    key: "action",
    dataIndex: "action"
  }];

  var _SWR = Object(swr__WEBPACK_IMPORTED_MODULE_2__["default"])("{\n      files(\n        hash: \"".concat(props.hash, "\"\n      ){\n        name\n        url\n      }\n    }"), _lib_graphql__WEBPACK_IMPORTED_MODULE_3__["default"]),
      data = _SWR.data,
      error = _SWR.error;

  console.log(error);

  if (!data) {
    return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Skeleton"], {}, void 0, false, {
      fileName: _jsxFileName,
      lineNumber: 39,
      columnNumber: 12
    }, this);
  }

  var dataSource = data.files.map(function (v) {
    v["filename"] = v.name;
    v["action"] = /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])("a", {
      href: v.url,
      target: "_blank_",
      children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Button"], {
        type: "primary",
        children: "Download"
      }, void 0, false, {
        fileName: _jsxFileName,
        lineNumber: 46,
        columnNumber: 9
      }, _this)
    }, void 0, false, {
      fileName: _jsxFileName,
      lineNumber: 45,
      columnNumber: 7
    }, _this);
    return v;
  });
  return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Table"], {
    columns: columns,
    dataSource: dataSource
  }, void 0, false, {
    fileName: _jsxFileName,
    lineNumber: 52,
    columnNumber: 10
  }, this);
}
_c = Filelist;

var _c;

$RefreshReg$(_c, "Filelist");

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
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vQ29tcG9uZW50L2ZpbGVsaXN0LnRzeCJdLCJuYW1lcyI6WyJGaWxlbGlzdCIsInByb3BzIiwiY29uc29sZSIsImxvZyIsImNvbHVtbnMiLCJ0aXRsZSIsImtleSIsImRhdGFJbmRleCIsIlNXUiIsImhhc2giLCJncmFwaHFsIiwiZGF0YSIsImVycm9yIiwiZGF0YVNvdXJjZSIsImZpbGVzIiwibWFwIiwidiIsIm5hbWUiLCJ1cmwiXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUNBO0FBQ0E7QUFNZSxTQUFTQSxRQUFULENBQWtCQyxLQUFsQixFQUF3QztBQUFBOztBQUNyREMsU0FBTyxDQUFDQyxHQUFSLENBQVlGLEtBQVo7QUFDQSxNQUFNRyxPQUFPLEdBQUcsQ0FDZDtBQUNFQyxTQUFLLEVBQUUsVUFEVDtBQUVFQyxPQUFHLEVBQUUsVUFGUDtBQUdFQyxhQUFTLEVBQUU7QUFIYixHQURjLEVBTWQ7QUFDRUYsU0FBSyxFQUFFLFFBRFQ7QUFFRUMsT0FBRyxFQUFFLFFBRlA7QUFHRUMsYUFBUyxFQUFFO0FBSGIsR0FOYyxDQUFoQjs7QUFGcUQsYUFlN0JDLG1EQUFHLDRDQUdaUCxLQUFLLENBQUNRLElBSE0sOERBU3pCQyxvREFUeUIsQ0FmMEI7QUFBQSxNQWU3Q0MsSUFmNkMsUUFlN0NBLElBZjZDO0FBQUEsTUFldkNDLEtBZnVDLFFBZXZDQSxLQWZ1Qzs7QUEyQnJEVixTQUFPLENBQUNDLEdBQVIsQ0FBWVMsS0FBWjs7QUFFQSxNQUFJLENBQUNELElBQUwsRUFBVztBQUNULHdCQUFPLHFFQUFDLDZDQUFEO0FBQUE7QUFBQTtBQUFBO0FBQUEsWUFBUDtBQUNEOztBQUVELE1BQUlFLFVBQVUsR0FBR0YsSUFBSSxDQUFDRyxLQUFMLENBQVdDLEdBQVgsQ0FBZSxVQUFDQyxDQUFELEVBQU87QUFDckNBLEtBQUMsQ0FBQyxVQUFELENBQUQsR0FBZ0JBLENBQUMsQ0FBQ0MsSUFBbEI7QUFDQUQsS0FBQyxDQUFDLFFBQUQsQ0FBRCxnQkFDRTtBQUFHLFVBQUksRUFBRUEsQ0FBQyxDQUFDRSxHQUFYO0FBQWdCLFlBQU0sRUFBQyxTQUF2QjtBQUFBLDZCQUNFLHFFQUFDLDJDQUFEO0FBQVEsWUFBSSxFQUFDLFNBQWI7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFERjtBQUFBO0FBQUE7QUFBQTtBQUFBLGFBREY7QUFLQSxXQUFPRixDQUFQO0FBQ0QsR0FSZ0IsQ0FBakI7QUFVQSxzQkFBTyxxRUFBQywwQ0FBRDtBQUFPLFdBQU8sRUFBRVosT0FBaEI7QUFBeUIsY0FBVSxFQUFFUztBQUFyQztBQUFBO0FBQUE7QUFBQTtBQUFBLFVBQVA7QUFDRDtLQTVDdUJiLFEiLCJmaWxlIjoic3RhdGljL3dlYnBhY2svcGFnZXMvdG9ycmVudC45YWUwYTA2MWM4ZjcyN2Q5ZGRhZS5ob3QtdXBkYXRlLmpzIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgVGFibGUsIFNrZWxldG9uLCBCdXR0b24gfSBmcm9tIFwiYW50ZFwiO1xuaW1wb3J0IFNXUiBmcm9tIFwic3dyXCI7XG5pbXBvcnQgZ3JhcGhxbCBmcm9tIFwiLi4vbGliL2dyYXBocWxcIjtcblxuaW50ZXJmYWNlIEZpbGVsaXN0UHJvcHMge1xuICBoYXNoOiBzdHJpbmc7XG59XG5cbmV4cG9ydCBkZWZhdWx0IGZ1bmN0aW9uIEZpbGVsaXN0KHByb3BzOiBGaWxlbGlzdFByb3BzKSB7XG4gIGNvbnNvbGUubG9nKHByb3BzKTtcbiAgY29uc3QgY29sdW1ucyA9IFtcbiAgICB7XG4gICAgICB0aXRsZTogXCJGaWxlbmFtZVwiLFxuICAgICAga2V5OiBcImZpbGVuYW1lXCIsXG4gICAgICBkYXRhSW5kZXg6IFwiZmlsZW5hbWVcIixcbiAgICB9LFxuICAgIHtcbiAgICAgIHRpdGxlOiBcIkFjdGlvblwiLFxuICAgICAga2V5OiBcImFjdGlvblwiLFxuICAgICAgZGF0YUluZGV4OiBcImFjdGlvblwiLFxuICAgIH0sXG4gIF07XG5cbiAgY29uc3QgeyBkYXRhLCBlcnJvciB9ID0gU1dSKFxuICAgIGB7XG4gICAgICBmaWxlcyhcbiAgICAgICAgaGFzaDogXCIke3Byb3BzLmhhc2h9XCJcbiAgICAgICl7XG4gICAgICAgIG5hbWVcbiAgICAgICAgdXJsXG4gICAgICB9XG4gICAgfWAsXG4gICAgZ3JhcGhxbFxuICApO1xuXG4gIGNvbnNvbGUubG9nKGVycm9yKTtcblxuICBpZiAoIWRhdGEpIHtcbiAgICByZXR1cm4gPFNrZWxldG9uIC8+O1xuICB9XG5cbiAgdmFyIGRhdGFTb3VyY2UgPSBkYXRhLmZpbGVzLm1hcCgodikgPT4ge1xuICAgIHZbXCJmaWxlbmFtZVwiXSA9IHYubmFtZTtcbiAgICB2W1wiYWN0aW9uXCJdID0gKFxuICAgICAgPGEgaHJlZj17di51cmx9IHRhcmdldD1cIl9ibGFua19cIj5cbiAgICAgICAgPEJ1dHRvbiB0eXBlPVwicHJpbWFyeVwiPkRvd25sb2FkPC9CdXR0b24+XG4gICAgICA8L2E+XG4gICAgKTtcbiAgICByZXR1cm4gdjtcbiAgfSk7XG5cbiAgcmV0dXJuIDxUYWJsZSBjb2x1bW5zPXtjb2x1bW5zfSBkYXRhU291cmNlPXtkYXRhU291cmNlfSAvPjtcbn1cbiJdLCJzb3VyY2VSb290IjoiIn0=
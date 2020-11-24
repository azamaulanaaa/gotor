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
  console.log(porps);
  var columns = [{
    title: "Filename",
    key: "filename",
    dataIndex: "filename"
  }, {
    title: "Action",
    key: "action",
    dataIndex: "action"
  }];

  var _SWR = Object(swr__WEBPACK_IMPORTED_MODULE_2__["default"])("{\n      files(\n        hash: \"".concat(props.hash, "}\"\n      ){\n        name\n        url\n      }\n    }"), _lib_graphql__WEBPACK_IMPORTED_MODULE_3__["default"]),
      data = _SWR.data,
      error = _SWR.error;

  if (!data) {
    return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Skeleton"], {}, void 0, false, {
      fileName: _jsxFileName,
      lineNumber: 37,
      columnNumber: 12
    }, this);
  }

  var dataSource = data.files.map(function (v) {
    v["filename"] = v.name;
    v["action"] = "";
    return v;
  });
  console.log(error);
  console.log(dataSource);
  return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__["jsxDEV"])(antd__WEBPACK_IMPORTED_MODULE_1__["Table"], {
    columns: columns,
    dataSource: dataSource
  }, void 0, false, {
    fileName: _jsxFileName,
    lineNumber: 49,
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
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vQ29tcG9uZW50L2ZpbGVsaXN0LnRzeCJdLCJuYW1lcyI6WyJGaWxlbGlzdCIsInByb3BzIiwiY29uc29sZSIsImxvZyIsInBvcnBzIiwiY29sdW1ucyIsInRpdGxlIiwia2V5IiwiZGF0YUluZGV4IiwiU1dSIiwiaGFzaCIsImdyYXBocWwiLCJkYXRhIiwiZXJyb3IiLCJkYXRhU291cmNlIiwiZmlsZXMiLCJtYXAiLCJ2IiwibmFtZSJdLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBO0FBQ0E7QUFDQTtBQU1lLFNBQVNBLFFBQVQsQ0FBa0JDLEtBQWxCLEVBQXdDO0FBQ3JEQyxTQUFPLENBQUNDLEdBQVIsQ0FBWUMsS0FBWjtBQUNBLE1BQU1DLE9BQU8sR0FBRyxDQUNkO0FBQ0VDLFNBQUssRUFBRSxVQURUO0FBRUVDLE9BQUcsRUFBRSxVQUZQO0FBR0VDLGFBQVMsRUFBRTtBQUhiLEdBRGMsRUFNZDtBQUNFRixTQUFLLEVBQUUsUUFEVDtBQUVFQyxPQUFHLEVBQUUsUUFGUDtBQUdFQyxhQUFTLEVBQUU7QUFIYixHQU5jLENBQWhCOztBQUZxRCxhQWU3QkMsbURBQUcsNENBR1pSLEtBQUssQ0FBQ1MsSUFITSwrREFTekJDLG9EQVR5QixDQWYwQjtBQUFBLE1BZTdDQyxJQWY2QyxRQWU3Q0EsSUFmNkM7QUFBQSxNQWV2Q0MsS0FmdUMsUUFldkNBLEtBZnVDOztBQTJCckQsTUFBSSxDQUFDRCxJQUFMLEVBQVc7QUFDVCx3QkFBTyxxRUFBQyw2Q0FBRDtBQUFBO0FBQUE7QUFBQTtBQUFBLFlBQVA7QUFDRDs7QUFFRCxNQUFJRSxVQUFVLEdBQUdGLElBQUksQ0FBQ0csS0FBTCxDQUFXQyxHQUFYLENBQWUsVUFBQ0MsQ0FBRCxFQUFPO0FBQ3JDQSxLQUFDLENBQUMsVUFBRCxDQUFELEdBQWdCQSxDQUFDLENBQUNDLElBQWxCO0FBQ0FELEtBQUMsQ0FBQyxRQUFELENBQUQsR0FBYyxFQUFkO0FBQ0EsV0FBT0EsQ0FBUDtBQUNELEdBSmdCLENBQWpCO0FBTUFmLFNBQU8sQ0FBQ0MsR0FBUixDQUFZVSxLQUFaO0FBQ0FYLFNBQU8sQ0FBQ0MsR0FBUixDQUFZVyxVQUFaO0FBRUEsc0JBQU8scUVBQUMsMENBQUQ7QUFBTyxXQUFPLEVBQUVULE9BQWhCO0FBQXlCLGNBQVUsRUFBRVM7QUFBckM7QUFBQTtBQUFBO0FBQUE7QUFBQSxVQUFQO0FBQ0Q7S0F6Q3VCZCxRIiwiZmlsZSI6InN0YXRpYy93ZWJwYWNrL3BhZ2VzL3RvcnJlbnQuOWViYWI5ZjRjYmZkZTJmMmE1MTMuaG90LXVwZGF0ZS5qcyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IFRhYmxlLCBTa2VsZXRvbiB9IGZyb20gXCJhbnRkXCI7XG5pbXBvcnQgU1dSIGZyb20gXCJzd3JcIjtcbmltcG9ydCBncmFwaHFsIGZyb20gXCIuLi9saWIvZ3JhcGhxbFwiO1xuXG5pbnRlcmZhY2UgRmlsZWxpc3RQcm9wcyB7XG4gIGhhc2g6IHN0cmluZztcbn1cblxuZXhwb3J0IGRlZmF1bHQgZnVuY3Rpb24gRmlsZWxpc3QocHJvcHM6IEZpbGVsaXN0UHJvcHMpIHtcbiAgY29uc29sZS5sb2cocG9ycHMpO1xuICBjb25zdCBjb2x1bW5zID0gW1xuICAgIHtcbiAgICAgIHRpdGxlOiBcIkZpbGVuYW1lXCIsXG4gICAgICBrZXk6IFwiZmlsZW5hbWVcIixcbiAgICAgIGRhdGFJbmRleDogXCJmaWxlbmFtZVwiLFxuICAgIH0sXG4gICAge1xuICAgICAgdGl0bGU6IFwiQWN0aW9uXCIsXG4gICAgICBrZXk6IFwiYWN0aW9uXCIsXG4gICAgICBkYXRhSW5kZXg6IFwiYWN0aW9uXCIsXG4gICAgfSxcbiAgXTtcblxuICBjb25zdCB7IGRhdGEsIGVycm9yIH0gPSBTV1IoXG4gICAgYHtcbiAgICAgIGZpbGVzKFxuICAgICAgICBoYXNoOiBcIiR7cHJvcHMuaGFzaH19XCJcbiAgICAgICl7XG4gICAgICAgIG5hbWVcbiAgICAgICAgdXJsXG4gICAgICB9XG4gICAgfWAsXG4gICAgZ3JhcGhxbFxuICApO1xuXG4gIGlmICghZGF0YSkge1xuICAgIHJldHVybiA8U2tlbGV0b24gLz47XG4gIH1cblxuICB2YXIgZGF0YVNvdXJjZSA9IGRhdGEuZmlsZXMubWFwKCh2KSA9PiB7XG4gICAgdltcImZpbGVuYW1lXCJdID0gdi5uYW1lO1xuICAgIHZbXCJhY3Rpb25cIl0gPSBcIlwiO1xuICAgIHJldHVybiB2O1xuICB9KTtcblxuICBjb25zb2xlLmxvZyhlcnJvcik7XG4gIGNvbnNvbGUubG9nKGRhdGFTb3VyY2UpO1xuXG4gIHJldHVybiA8VGFibGUgY29sdW1ucz17Y29sdW1uc30gZGF0YVNvdXJjZT17ZGF0YVNvdXJjZX0gLz47XG59XG4iXSwic291cmNlUm9vdCI6IiJ9
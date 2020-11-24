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
      lineNumber: 36,
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
    lineNumber: 48,
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
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vQ29tcG9uZW50L2ZpbGVsaXN0LnRzeCJdLCJuYW1lcyI6WyJGaWxlbGlzdCIsInByb3BzIiwiY29sdW1ucyIsInRpdGxlIiwia2V5IiwiZGF0YUluZGV4IiwiU1dSIiwiaGFzaCIsImdyYXBocWwiLCJkYXRhIiwiZXJyb3IiLCJkYXRhU291cmNlIiwiZmlsZXMiLCJtYXAiLCJ2IiwibmFtZSIsImNvbnNvbGUiLCJsb2ciXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUNBO0FBQ0E7QUFNZSxTQUFTQSxRQUFULENBQWtCQyxLQUFsQixFQUF3QztBQUNyRCxNQUFNQyxPQUFPLEdBQUcsQ0FDZDtBQUNFQyxTQUFLLEVBQUUsVUFEVDtBQUVFQyxPQUFHLEVBQUUsVUFGUDtBQUdFQyxhQUFTLEVBQUU7QUFIYixHQURjLEVBTWQ7QUFDRUYsU0FBSyxFQUFFLFFBRFQ7QUFFRUMsT0FBRyxFQUFFLFFBRlA7QUFHRUMsYUFBUyxFQUFFO0FBSGIsR0FOYyxDQUFoQjs7QUFEcUQsYUFjN0JDLG1EQUFHLDRDQUdaTCxLQUFLLENBQUNNLElBSE0sK0RBU3pCQyxvREFUeUIsQ0FkMEI7QUFBQSxNQWM3Q0MsSUFkNkMsUUFjN0NBLElBZDZDO0FBQUEsTUFjdkNDLEtBZHVDLFFBY3ZDQSxLQWR1Qzs7QUEwQnJELE1BQUksQ0FBQ0QsSUFBTCxFQUFXO0FBQ1Qsd0JBQU8scUVBQUMsNkNBQUQ7QUFBQTtBQUFBO0FBQUE7QUFBQSxZQUFQO0FBQ0Q7O0FBRUQsTUFBSUUsVUFBVSxHQUFHRixJQUFJLENBQUNHLEtBQUwsQ0FBV0MsR0FBWCxDQUFlLFVBQUNDLENBQUQsRUFBTztBQUNyQ0EsS0FBQyxDQUFDLFVBQUQsQ0FBRCxHQUFnQkEsQ0FBQyxDQUFDQyxJQUFsQjtBQUNBRCxLQUFDLENBQUMsUUFBRCxDQUFELEdBQWMsRUFBZDtBQUNBLFdBQU9BLENBQVA7QUFDRCxHQUpnQixDQUFqQjtBQU1BRSxTQUFPLENBQUNDLEdBQVIsQ0FBWVAsS0FBWjtBQUNBTSxTQUFPLENBQUNDLEdBQVIsQ0FBWU4sVUFBWjtBQUVBLHNCQUFPLHFFQUFDLDBDQUFEO0FBQU8sV0FBTyxFQUFFVCxPQUFoQjtBQUF5QixjQUFVLEVBQUVTO0FBQXJDO0FBQUE7QUFBQTtBQUFBO0FBQUEsVUFBUDtBQUNEO0tBeEN1QlgsUSIsImZpbGUiOiJzdGF0aWMvd2VicGFjay9wYWdlcy90b3JyZW50LjFhNDRmYzg0YTNmNDdiOTJiNTg1LmhvdC11cGRhdGUuanMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBUYWJsZSwgU2tlbGV0b24gfSBmcm9tIFwiYW50ZFwiO1xuaW1wb3J0IFNXUiBmcm9tIFwic3dyXCI7XG5pbXBvcnQgZ3JhcGhxbCBmcm9tIFwiLi4vbGliL2dyYXBocWxcIjtcblxuaW50ZXJmYWNlIEZpbGVsaXN0UHJvcHMge1xuICBoYXNoOiBzdHJpbmc7XG59XG5cbmV4cG9ydCBkZWZhdWx0IGZ1bmN0aW9uIEZpbGVsaXN0KHByb3BzOiBGaWxlbGlzdFByb3BzKSB7XG4gIGNvbnN0IGNvbHVtbnMgPSBbXG4gICAge1xuICAgICAgdGl0bGU6IFwiRmlsZW5hbWVcIixcbiAgICAgIGtleTogXCJmaWxlbmFtZVwiLFxuICAgICAgZGF0YUluZGV4OiBcImZpbGVuYW1lXCIsXG4gICAgfSxcbiAgICB7XG4gICAgICB0aXRsZTogXCJBY3Rpb25cIixcbiAgICAgIGtleTogXCJhY3Rpb25cIixcbiAgICAgIGRhdGFJbmRleDogXCJhY3Rpb25cIixcbiAgICB9LFxuICBdO1xuXG4gIGNvbnN0IHsgZGF0YSwgZXJyb3IgfSA9IFNXUihcbiAgICBge1xuICAgICAgZmlsZXMoXG4gICAgICAgIGhhc2g6IFwiJHtwcm9wcy5oYXNofX1cIlxuICAgICAgKXtcbiAgICAgICAgbmFtZVxuICAgICAgICB1cmxcbiAgICAgIH1cbiAgICB9YCxcbiAgICBncmFwaHFsXG4gICk7XG5cbiAgaWYgKCFkYXRhKSB7XG4gICAgcmV0dXJuIDxTa2VsZXRvbiAvPjtcbiAgfVxuXG4gIHZhciBkYXRhU291cmNlID0gZGF0YS5maWxlcy5tYXAoKHYpID0+IHtcbiAgICB2W1wiZmlsZW5hbWVcIl0gPSB2Lm5hbWU7XG4gICAgdltcImFjdGlvblwiXSA9IFwiXCI7XG4gICAgcmV0dXJuIHY7XG4gIH0pO1xuXG4gIGNvbnNvbGUubG9nKGVycm9yKTtcbiAgY29uc29sZS5sb2coZGF0YVNvdXJjZSk7XG5cbiAgcmV0dXJuIDxUYWJsZSBjb2x1bW5zPXtjb2x1bW5zfSBkYXRhU291cmNlPXtkYXRhU291cmNlfSAvPjtcbn1cbiJdLCJzb3VyY2VSb290IjoiIn0=
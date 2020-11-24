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
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vQ29tcG9uZW50L2ZpbGVsaXN0LnRzeCJdLCJuYW1lcyI6WyJGaWxlbGlzdCIsInByb3BzIiwiY29uc29sZSIsImxvZyIsImNvbHVtbnMiLCJ0aXRsZSIsImtleSIsImRhdGFJbmRleCIsIlNXUiIsImhhc2giLCJncmFwaHFsIiwiZGF0YSIsImVycm9yIiwiZGF0YVNvdXJjZSIsImZpbGVzIiwibWFwIiwidiIsIm5hbWUiXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUNBO0FBQ0E7QUFNZSxTQUFTQSxRQUFULENBQWtCQyxLQUFsQixFQUF3QztBQUNyREMsU0FBTyxDQUFDQyxHQUFSLENBQVlGLEtBQVo7QUFDQSxNQUFNRyxPQUFPLEdBQUcsQ0FDZDtBQUNFQyxTQUFLLEVBQUUsVUFEVDtBQUVFQyxPQUFHLEVBQUUsVUFGUDtBQUdFQyxhQUFTLEVBQUU7QUFIYixHQURjLEVBTWQ7QUFDRUYsU0FBSyxFQUFFLFFBRFQ7QUFFRUMsT0FBRyxFQUFFLFFBRlA7QUFHRUMsYUFBUyxFQUFFO0FBSGIsR0FOYyxDQUFoQjs7QUFGcUQsYUFlN0JDLG1EQUFHLDRDQUdaUCxLQUFLLENBQUNRLElBSE0sK0RBU3pCQyxvREFUeUIsQ0FmMEI7QUFBQSxNQWU3Q0MsSUFmNkMsUUFlN0NBLElBZjZDO0FBQUEsTUFldkNDLEtBZnVDLFFBZXZDQSxLQWZ1Qzs7QUEyQnJELE1BQUksQ0FBQ0QsSUFBTCxFQUFXO0FBQ1Qsd0JBQU8scUVBQUMsNkNBQUQ7QUFBQTtBQUFBO0FBQUE7QUFBQSxZQUFQO0FBQ0Q7O0FBRUQsTUFBSUUsVUFBVSxHQUFHRixJQUFJLENBQUNHLEtBQUwsQ0FBV0MsR0FBWCxDQUFlLFVBQUNDLENBQUQsRUFBTztBQUNyQ0EsS0FBQyxDQUFDLFVBQUQsQ0FBRCxHQUFnQkEsQ0FBQyxDQUFDQyxJQUFsQjtBQUNBRCxLQUFDLENBQUMsUUFBRCxDQUFELEdBQWMsRUFBZDtBQUNBLFdBQU9BLENBQVA7QUFDRCxHQUpnQixDQUFqQjtBQU1BZCxTQUFPLENBQUNDLEdBQVIsQ0FBWVMsS0FBWjtBQUNBVixTQUFPLENBQUNDLEdBQVIsQ0FBWVUsVUFBWjtBQUVBLHNCQUFPLHFFQUFDLDBDQUFEO0FBQU8sV0FBTyxFQUFFVCxPQUFoQjtBQUF5QixjQUFVLEVBQUVTO0FBQXJDO0FBQUE7QUFBQTtBQUFBO0FBQUEsVUFBUDtBQUNEO0tBekN1QmIsUSIsImZpbGUiOiJzdGF0aWMvd2VicGFjay9wYWdlcy90b3JyZW50LjI1OTVjZTIzYjVlYjU3OTM3MThjLmhvdC11cGRhdGUuanMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBUYWJsZSwgU2tlbGV0b24gfSBmcm9tIFwiYW50ZFwiO1xuaW1wb3J0IFNXUiBmcm9tIFwic3dyXCI7XG5pbXBvcnQgZ3JhcGhxbCBmcm9tIFwiLi4vbGliL2dyYXBocWxcIjtcblxuaW50ZXJmYWNlIEZpbGVsaXN0UHJvcHMge1xuICBoYXNoOiBzdHJpbmc7XG59XG5cbmV4cG9ydCBkZWZhdWx0IGZ1bmN0aW9uIEZpbGVsaXN0KHByb3BzOiBGaWxlbGlzdFByb3BzKSB7XG4gIGNvbnNvbGUubG9nKHByb3BzKTtcbiAgY29uc3QgY29sdW1ucyA9IFtcbiAgICB7XG4gICAgICB0aXRsZTogXCJGaWxlbmFtZVwiLFxuICAgICAga2V5OiBcImZpbGVuYW1lXCIsXG4gICAgICBkYXRhSW5kZXg6IFwiZmlsZW5hbWVcIixcbiAgICB9LFxuICAgIHtcbiAgICAgIHRpdGxlOiBcIkFjdGlvblwiLFxuICAgICAga2V5OiBcImFjdGlvblwiLFxuICAgICAgZGF0YUluZGV4OiBcImFjdGlvblwiLFxuICAgIH0sXG4gIF07XG5cbiAgY29uc3QgeyBkYXRhLCBlcnJvciB9ID0gU1dSKFxuICAgIGB7XG4gICAgICBmaWxlcyhcbiAgICAgICAgaGFzaDogXCIke3Byb3BzLmhhc2h9fVwiXG4gICAgICApe1xuICAgICAgICBuYW1lXG4gICAgICAgIHVybFxuICAgICAgfVxuICAgIH1gLFxuICAgIGdyYXBocWxcbiAgKTtcblxuICBpZiAoIWRhdGEpIHtcbiAgICByZXR1cm4gPFNrZWxldG9uIC8+O1xuICB9XG5cbiAgdmFyIGRhdGFTb3VyY2UgPSBkYXRhLmZpbGVzLm1hcCgodikgPT4ge1xuICAgIHZbXCJmaWxlbmFtZVwiXSA9IHYubmFtZTtcbiAgICB2W1wiYWN0aW9uXCJdID0gXCJcIjtcbiAgICByZXR1cm4gdjtcbiAgfSk7XG5cbiAgY29uc29sZS5sb2coZXJyb3IpO1xuICBjb25zb2xlLmxvZyhkYXRhU291cmNlKTtcblxuICByZXR1cm4gPFRhYmxlIGNvbHVtbnM9e2NvbHVtbnN9IGRhdGFTb3VyY2U9e2RhdGFTb3VyY2V9IC8+O1xufVxuIl0sInNvdXJjZVJvb3QiOiIifQ==
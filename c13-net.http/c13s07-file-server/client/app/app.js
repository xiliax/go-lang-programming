(function (ng) {
    ng.module('MyApp', [])
        .controller('MainController', function () {
            var self = this;
            self.name = "Some Name";
        })
})(angular);
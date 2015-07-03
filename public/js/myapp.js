(function() {
	var app = angular.module('myapp', []);
	app.controller('MainController', function() {
		this.products = products;
		this.items = [];

		this.addItem = function(index, i) {
			quantity = 1;
			for (var i = this.items.length - 1; i >= 0; i--) {
				item = this.items[i];
				if (item.item.name === products[index].name) {
					quantity = item.qty + 1;
					if (quantity === 0) {
						quantity = 1;
					};
					this.items.splice(i, 1);
					break;
				};
			};
			this.items.push({item:products[index], qty:quantity});
		};

		this.consume = function() {
			this.items = [];
		};

		this.charge = function() {
			this.items = [];
		};

	});

	app.controller('TabController', function(){
		this.current = 1;
		this.setTab = function(tab) {
			this.current = tab;
		};
		this.isTab = function(tab) {
			if (this.current === tab) {
				return true
			};
			return false;
		};
	});

	var products = [
		{
			name: '250ml Milk',
			image : '/img/naiping.png'
		},
		{
			name: '500ml Milk',
			image : '/img/naiping.png'
		},
		{
			name: 'Yugart',
			image : '/img/naiping.png'
		},
		{
			name: 'Pudding',
			image : '/img/naiping.png'
		},
		{
			name: 'Shuangpi Milk',
			image : '/img/naiping.png'
		},
		{
			name: 'Bread',
			image : '/img/naiping.png'
		},
		{
			name: '250ml Milk * 2',
			image : '/img/naiping.png'
		}
	];
})();
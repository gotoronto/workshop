/*jshint eqeqeq:false */
(function (window) {
  'use strict';

  /**
   * Creates a new client side storage object and will create an empty
   * collection if no collection already exists.
   *
   * @param {string} name The name of our DB we want to use
   * @param {function} callback Our fake DB uses callbacks because in
   * real life you probably would be making AJAX calls
   */
  function Store() { }

  /**
   * Finds items based on a query given as a JS object
   *
   * @param {object} query The query to match against (i.e. {foo: 'bar'})
   * @param {function} callback  The callback to fire when the query has
   * completed running
   *
   * @example
   * db.find({foo: 'bar', hello: 'world'}, function (data) {
   *   // data will return any items that have foo: bar and
   *   // hello: world in their properties
   * });
   */
  Store.prototype.find = function (id, callback) {
    fetch(`/todos/${id}`).then((resp) => {
      resp.json().then((todo) => {
        callback.call(this, [todo]);
      });
    });
  };

  Store.prototype.query = function (query, callback) {
    this.findAll((todos) => {
      callback.call(this, todos.filter(function (todo) {
        for (var q in query) {
          if (query[q] !== todo[q]) {
            return false;
          }
        }
        return true;
      }));
    })
  };

  /**
   * Will retrieve all data from the collection
   *
   * @param {function} callback The callback to fire upon retrieving data
   */
  Store.prototype.findAll = function (callback) {
    callback = callback || function () {};
    fetch("/todos").then((data) => {
      data.json().then((todos) => {
        callback.call(this, todos);
      });
    });
  };

  /**
   * Will save the given data to the DB. If no item exists it will create a new
   * item, otherwise it'll simply update an existing item's properties
   *
   * @param {object} updateData The data to save back into the DB
   * @param {function} callback The callback to fire after saving
   * @param {number} id An optional param to enter an ID of an item to update
   */
  Store.prototype.save = function (updateData, callback, id) {
    var url = "/todos"
    if (id) {
      url += `/${id}`
    }
    fetch(url, {
      method: 'POST',
      body: JSON.stringify(updateData)
    }).then((resp) => {
      resp.json().then((todos) => {
        callback.call(this, todos);
      });
    });
  };

  /**
   * Will remove an item from the Store based on its ID
   *
   * @param {number} id The ID of the item you want to remove
   * @param {function} callback The callback to fire after saving
   */
  Store.prototype.remove = function (id, callback) {
    fetch(`/todos/${id}/delete`, {
      method: 'POST',
    }).then((resp) => {
      resp.json().then((todos) => {
        callback.call(this, todos);
      });
    });
  };

  /**
   * Will drop all storage and start fresh
   *
   * @param {function} callback The callback to fire after dropping the data
   */
  Store.prototype.drop = function (callback) {
    this.findAll((todos) => {
      todos.forEach((todo) => {
        this.remove(todo.id)
      });
      callback.call(this);
    });
  };

  // Export to window
  window.app = window.app || {};
  window.app.Store = Store;
})(window);

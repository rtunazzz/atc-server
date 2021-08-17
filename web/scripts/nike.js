/**
 * Finds a cookie by name and returns it's value.
 * @param {String} name Name of the cookie
 * @returns {String|void} Value of the cookie if found, otherwise void
 */
function getCookieValue(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(";").shift();
}

/**
 * Generates an UUID.
 * (UUID should be unique and this does not guarantee the uniques. That's why it's unsafe.)
 * @returns {String} the generated UUID
 */
function createUnsafeUUID() {
  let dt = new Date().getTime();
  const uuid = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(
    /[xy]/g,
    function (c) {
      const r = (dt + Math.random() * 16) % 16 | 0;
      dt = Math.floor(dt / 16);
      return (c == "x" ? r : (r & 0x3) | 0x8).toString(16);
    }
  );
  return uuid;
}

/**
 * Fetches the default wishlist ID
 * @param {String} bearer Auth bearer token from https://unite.nike.com/auth/unite_session_cookies/v1
 * @param {String} country country of the wishlist
 * @returns {Promise<String|void>} default wishlist ID on success, undefined otherwise
 */
async function getDefaultWishlistId(bearer, country) {
  const url = `https://api.nike.com/buy/lists/v1?filter=isDefault(true)&filter=country(${country.toUpperCase()})`;

  const res = await fetch(url, {
    method: "GET",
    headers: {
      "content-type": "application/json; charset=UTF-8",
      authorization: `Bearer ${bearer}`,
    },
  });

  console.log(res);
  const j = await res.json();
  console.log(j);

  return j?.objects?.[0]?.id;
}

/**
 * Adds the productId passed in into the wishlist specified by its ID.
 * @param {String} bearer Auth bearer token from 'https://unite.nike.com/auth/unite_session_cookies/v1'
 * @param {String} country country of the wishlist
 * @param {String} wishlistId ID of the wishlist to add the product to
 * @param {String} productId ID of the product to add to wihslist
 * @returns {Promise<Boolean>} true on success, false otherwise
 */
async function addToWishlist(bearer, country, wishlistId, productId) {
  const url = `https://api.nike.com/buy/list_items/v1/${createUnsafeUUID()}`;

  const data = {
    id: productId,
    productId,
    skuId: null,
    wishlistId,
  };

  const res = await fetch(url, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: {
      "content-type": "application/json; charset=UTF-8",
      authorization: `Bearer ${bearer}`,
    },
  });
  console.log(res);

  return res.status == 200;
}

/**
 * Handles adding a product to a wishlist
 * @returns {Promise<void>}
 */
async function handleWishlist() {
  const token = getCookieValue("nike_auth");
  const country = getCookieValue("nike_locale");

  console.log("COUNTRY:", country);
  console.log("TOKEN:", token);

  if (!token || !country) {
    location.replace("/nike/set_locale");
    return;
  }

  const productId = document.getElementById("productId")?.value;
  if (!productId) {
    console.error("No product ID found!");
    location.replace(`https://nike.com/${country}/cart`);
    return;
  }

  try {
    const wishlistId = await getDefaultWishlistId(token, country);
    if (!wishlistId) {
      location.replace("/nike/set_locale");
      return;
    }
    const added = await addToWishlist(token, country, wishlistId, productId);

    console.log(added);
    location.replace(`https://nike.com/${country}/cart`);
  } catch (err) {
    console.log("Failed to handle wishlist!");
    console.error(err);
    location.replace("/nike/set_locale");
  }
}

handleWishlist();

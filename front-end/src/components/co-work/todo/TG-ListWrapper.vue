<template>
    <div class="boardLists">
        <div class="listWrapper">
            <div class="listContent">
                <div class="listTitle">
                    <div class="listTitleControllerBox">
                        <span class="listTitleControllerA" v-show="titleController">
                            <h3 class="listContentTitle"> Todo </h3>
                        </span>
                        <span class="listTitleControllerB" v-show="!titleController">
                            <textarea class="listTitleChange" ref="focusValue"></textarea>
                        </span>
                    </div>
                    <span>
                        <i class="fa-solid fa-ellipsis"></i>
                    </span>
                </div>
                <div class="listCards">
                    <a class="listCardslist"> add your Todos... </a>
                </div>
                <div class="addACardController">
                    <div class="addACardControllerA" v-show="addCardController">
                        <a class="addACardButton">+ Add a card</a>
                    </div>
                    <div class="addACardControllerB" v-show="!addCardController">
                        <textarea class="addCardInCardsList" placeholder="Enter a title for this card..."></textarea>
                        <div class="addACardButtonBox">
                            <div class="addCardButton">
                                <input class="addCard" type="button" value="Add card">
                                <span class="closeAddCard">
                                    <i class="fa-solid fa-xmark"></i>
                                </span>
                            </div>
                            <span class="addCardOptions">
                                <i class="fa-solid fa-ellipsis"></i>
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            titleController: true,
            addCardController: true,
            titleValue: "",
            newCardValue: "",
        }
    }
}
</script>

<style>
/* Wrapper Style Sheet */
.boardLists {
    display: flex;
    flex-direction: row;
}

.listWrapper {
    flex-shrink: 0;
    width: 272px;
    max-height: max-content;
    margin-right: 10px;
}

.listContent {
    display: flex;
    flex-direction: column;
    background-color: #ebecf0;
    border-radius: 5px;
    padding-left: 10px;
    padding-bottom: 10px;
}

/* Title Style Sheet */
.listTitle {
    display: flex;
    flex-direction: row;
}

.listTitleControllerBox {
    flex-shrink: 0;
    width: 220px;
    margin-right: 10px;
}

.listContentTitle {
    margin: 8px 3px 8px 5px;
    color: grey;
}

.fa-ellipsis {
    padding-top: 10px;
    color: gray;
    margin-right: 15px;
    font-size: 20px;
}

.listTitleChange {
    border: 2px solid cornflowerblue;
    border-radius: 3px;
    width: 96%;
    height: 25px;
    margin: 5px 5px 0 0;
    padding: 0 5px 2px 5px;
    text-align: left;
    size: 15px;
    color: gray;
    font-family: sans-serif;
    font-size: 18px;
    font-weight: bold;
    resize: none;
}

/* Card Style Sheet */
.listCards {
    display: flex;
    flex-direction: column;
}

.listCardslist {
    display: block;
    background-color: white;
    border-radius: 3px;
    margin-bottom: 7px;
    width: 243px;
    min-height: 25px;
    box-shadow: 0px 2px 2px #E2E2E2;
    padding: 5px;
    min-height: 20px;
    color: grey;
    font-size: 15px;
}

/* Add Card Style Sheet */
.addACardController {
    display: flex;
    flex-direction: column;
}

.addACardControllerA {
    display: flex;
    flex-direction: space-between;
    max-height: 38px;
    min-height: 38px;
}


.addACardButton {
    background-color: #ebecf0;
    border-radius: 3px;
    display: block;
    color: #5e6c84;
    flex: 1 0 auto;
    margin-top: 2px;
    padding: 7px 8px 4px 0;
    position: relative;
    max-height: 27px;
    font-size: 15px;
}

.addACardControllerB {
    display: block;
    flex-shrink: 0;
}

.addCardInCardsList {
    border: 0 solid black;
    border-radius: 3px;
    width: 242px;
    min-height: 70px;
    box-shadow: 0px 2px 2px #ebecf0;
    margin-bottom: 3px;
    padding-left: 7px;
    color: gray;
    font-weight: 100;
    font-size: 15px;
    font-family: sans-serif;
    resize: none;
}

.addACardButtonBox {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.addCardButton {
    display: flex;
    flex-direction: row;
    max-height: 35px;

}

.addCard {
    border: 0 solid black;
    border-radius: 3px;
    background-color: cornflowerblue;
    box-shadow: 0px 2px 2px #ebecf0;
    color: white;
    flex-shrink: 0;
    width: 70px;
    height: 35px;
}

.closeAddCard {
    margin-left: 10px;
    margin-top: 8px;
}

.fa-xmark {
    color: grey;
    font-size: 20px;
}

.addCardOptions {
    margin-right: 10px;
}
</style>
<style scoped>
.fa-ellipsis {
    color: gray;
    font-size: 20px;
    padding-top: 8px;
    margin-right: 5px;
}
</style>
